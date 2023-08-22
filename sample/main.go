package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error is", err)
	}
	defer res.Body.Close()

	var body map[string]interface{}

	err1 := json.NewDecoder(res.Body).Decode(&body)
	if err1 != nil {
		fmt.Println("eE", err)

	}
	fmt.Println(body)
}
