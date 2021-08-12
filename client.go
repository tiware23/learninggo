package main

import (
	"crypto/tls"
	"fmt"
	"kubernetes/client"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	urlApi := "url"
	token := "<token>"

	postRequest := client.PostPod(urlApi, token)
	if postRequest == 201 {
		fmt.Printf("Created Pod status: %d\n", postRequest)
	} else if postRequest == 409 {
		fmt.Printf("The resource already exists status: %d\n", postRequest)
	} else {
		fmt.Println("Found error durind the process")
	}

	request := client.GetPod(urlApi, token)
	fmt.Println(request)

}
