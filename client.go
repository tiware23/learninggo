// Script to get and post Pod object from k8s api manager
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func httpMethod(method, uri, token string, body io.Reader) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, body)
	checkError(err)

	bearer := "Bearer " + token
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)
	checkError(err)

	return resp
}

func getPod(uri string, token string) {
	res := httpMethod("GET", uri, token, nil)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, body, "", "\t")
	log.Println(string(prettyJSON.Bytes()))
}

func postPod(uri string, token string) {
	body, err := os.Open("pod.json")
	res := httpMethod("POST", uri, token, body)
	checkError(err)
	defer body.Close()

	fmt.Println("Response status: ", res.Status)
	io.Copy(os.Stdout, res.Body)

}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	urlApi := "https://kubernetes.docker.internal:6443/api/v1/namespaces/default/pods"
	token := "<token>"

	getPod(urlApi, token)
	// postPod(urlApi, token)

}
