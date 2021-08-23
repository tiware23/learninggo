// Client script to get and post Pod object from k8s api manager
package client

import (
	"encoding/json"
	"fmt"
)

type Pod struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Property
}

type Property struct {
	Metadata map[string]interface{} `json:"metadata"`
	Spec     map[string]interface{} `json:"spec"`
}

// GeneratePod generates json Pod file
func GeneratePod(name, image string) []byte {
	// Property Metadata
	labels := make(map[string]string)
	labels["app"] = name

	var p Property

	p.Metadata = make(map[string]interface{})

	p.Metadata["name"] = name

	p.Metadata["labels"] = labels

	// Property Spec
	p.Spec = make(map[string]interface{})
	var c []map[string]string
	a1 := map[string]string{"name": name, "image": image}
	c = append(c, a1)

	p.Spec["containers"] = c

	u := Pod{Kind: "Pod", ApiVersion: "v1"}
	u.Property = p

	jU, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	return jU
}
