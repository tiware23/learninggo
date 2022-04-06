package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Adding Application's manifests in a Map of string:
	// Specifying XpCryptor Objects
	xpcryptorFiles := map[string]string{
		"configmaps": "xpcryptor-content",
		"secrets":    "xpcryptor-key",
	}
	// Specifying The remain Objects
	objectFiles := map[string]string{
		"configmaps":               "/api/v1/namespaces/",
		"deployments":              "/apis/apps/v1/namespaces/",
		"resourcequotas":           "/api/v1/namespaces/",
		"secrets":                  "/api/v1/namespaces/",
		"services":                 "/api/v1/namespaces/",
		"serviceaccounts":          "/api/v1/namespaces/",
		"ingresses":                "/apis/networking.k8s.io/v1/namespaces/",
		"rolebindings":             "/apis/rbac.authorization.k8s.io/v1/namespaces/",
		"roles":                    "/apis/rbac.authorization.k8s.io/v1/namespaces/",
		"horizontalpodautoscalers": "/apis/autoscaling/v1/namespaces/",
	}

	// This namespaces.txt file must contain all namepsaces which does not have
	// Helm's Labels and Annotations.
	n, err := readFile("namespaces.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range n {
		patchObjecFiles(objectFiles, i)
		patchXpcrytorFiles(xpcryptorFiles, i)
	}

}

// patchObjecFiles run the PATCH against to objectFiles
func patchObjecFiles(objectFiles map[string]string, namespace string) {
	for k, v := range objectFiles {
		resp, err := httpMethod("PATCH", "http://localhost:8001"+v+namespace+"/"+k+"/"+namespace, namespace)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Namespace %s - Object: %s - Patched: %s\n", namespace, k, resp.Status)
	}
}

// patchXpcrytorFiles run the PATCH against to XpCryptor files
func patchXpcrytorFiles(xpcryptorFiles map[string]string, namespace string) {
	for k, v := range xpcryptorFiles {
		resp, err := httpMethod("PATCH", "http://localhost:8001/api/v1/namespaces/"+namespace+"/"+k+"/"+v, namespace)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Namespace %s - Object: %s: %s - Patched: %s\n", namespace, k, resp.Status, v)
	}
}

// httpMethod was built in order to create a http new request with reusing tcp connections.
func httpMethod(method, uri, namespace string) (*http.Response, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	parsePatch := fmt.Sprintf(`{"metadata":{"labels":{"app.kubernetes.io/managed-by":"Helm"}}},
							{"annotations":{"meta.helm.sh/release-name":"%s","meta.helm.sh/release-namespace":"%s"}}}`, namespace, namespace)

	helmBuff := []byte(parsePatch)
	body := bytes.NewBuffer(helmBuff)

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", `application/strategic-merge-patch+json`)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	return resp, err
}

//  readFile read the namespaces.txt
func readFile(file string) ([]string, error) {
	var sValues []string

	f, err := os.Open(file)
	if err != nil {
		return sValues, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		sValues = append(sValues, strings.TrimSpace(scanner.Text()))
	}

	err = f.Close()
	if err != nil {
		return sValues, err
	}

	if scanner.Err() != nil {
		return sValues, scanner.Err()
	}

	return sValues, nil
}
