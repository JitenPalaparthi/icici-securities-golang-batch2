package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8083/users"
	method := "POST"

	payload := strings.NewReader(`{
    "name": "Priya",
    "email": "priya@outlook.com",
    "contact": "9618558500"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// write shell command to run the client 1000 times
// for i in {1..1000}; do go run client/client.go; done
