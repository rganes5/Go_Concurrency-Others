package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	Body   string `json:"body"`
}

func main() {
	//url := "https://jsonplaceholder.typicode.com/posts/1"
	url := "https://jsonplaceholder.typicode.com/posts"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	//var body map[string]interface{}
	var resp []Post
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		fmt.Println("failed", err)
	}

	fmt.Println("The data is :")
	for _, val := range resp {
		fmt.Println("User id", val.UserID)
		fmt.Println("body is ", val.Body)
	}
	//fmt.Println("User id:", resp.UserID)
	//fmt.Println("Body is:", resp.Body)

	//Ioutil method

	// Read the response body
	/*
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Unmarshal the JSON data
		//var posts Post
		var posts []Post
		err = json.Unmarshal(body, &posts)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Display userId and body for each post
		//fmt.Println(posts.Body)
		//fmt.Println(posts.UserID)
		for _, post := range posts {
			fmt.Printf("UserID: %d\n", post.UserID)
			fmt.Printf("Body: %s\n\n", post.Body)
		}
	*/
}
