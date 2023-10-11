package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message"`
}

type Body struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	//var response PingResponse
	//body.Message = "pong"
	response := PingResponse{Message: "pong"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	thirdPartyURL := "https://jsonplaceholder.typicode.com/posts/1"
	res, err := http.Get(thirdPartyURL)
	if err != nil {
		http.Error(w, "Failed to get the endpoint", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	var body Body

	//var body map[string]interface{}

	/*
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/

	err1 := json.NewDecoder(res.Body).Decode(&body)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "fucked", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func main() {

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/get", getHandler)
	fmt.Println("Server running on port :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start a server")
	}

}
