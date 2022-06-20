package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	token := "s.OyKNu4g5vyO3d2fqwioo"
	vaultAddr := "vault.example.com"

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(token)

	//writing the data
	inputData := map[string]interface{}{
		"data": map[string]interface{}{
			"first": "ankit",
		},
	}
	output, err := client.Logical().Write("secret/data/abd", inputData)
	fmt.Println(output)
	if err != nil {
		panic(err)
	}

	//deleting the data
	data, err := client.Logical().Read("secret/data/hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Data)

	//deleting the data
	output, err = client.Logical().Delete("secret/metadata/abd")
	fmt.Println(output)
	if err != nil {
		panic(err)
	}
}
