// Client script to get and post Pod object from k8s api manager
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// HttpMethod creates a NewRequests for http method
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

// GetPod gets information throught HttpMethod
func GetPod(uri string, token string) string {
	res := httpMethod("GET", uri, token, nil)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, body, "", "\t")
	return string(prettyJSON.Bytes())
}

// PostPod gets information throught HttpMethod
func PostPod(uri string, token string) int {
	body, err := os.Open("pod.json")
	checkError(err)
	res := httpMethod("POST", uri, token, body)
	defer body.Close()

	return res.StatusCode

}
