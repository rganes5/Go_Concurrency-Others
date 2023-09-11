package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Respond to a ping request with a pong response
	response := PingResponse{Message: "pong"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func fetchAndForwardHandler(w http.ResponseWriter, r *http.Request) {
	// Make an HTTP GET request to a third-party endpoint
	thirdPartyURL := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(thirdPartyURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Forward the response as is
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {
	// Create a simple HTTP server
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/fetch-and-forward", fetchAndForwardHandler)

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
