package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var _port = ":8000"

func main() {

	http.HandleFunc("/account", getDataHandler)
	log.Println("listening on ", _port)
	http.ListenAndServe(_port, nil)
}

type account struct {
	Name string
	Age  int
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {

	// check the request method manually
	if r.Method != http.MethodGet {
		return
	}

	// create a data
	a := account{
		Name: "My Name",
		Age:  25,
	}

	// marshal data to json
	data, err := json.Marshal(a)
	if err != nil {
		if err != nil {
			log.Println("Failed to unmarshal data to json: ", err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json") // optional
	// write data to the response
	_, err = w.Write(data)
	if err != nil {
		log.Println("Failed to write account details on response: ", err)
		return
	}
}
