package main

import (
	"encoding/json"
	"fmt"
)

// sample json object to be parsed
var companyJSON = `{
	"name" : "GOLinuxCloud",
        "years_of_service" : "5",
        "nature_of_company" : "Online Academy",
        "no_of_staff" : "10"
}`

func main() {
	ch := make(chan struct{})
	data := make(map[string]interface{})
	go func() {
		err := json.Unmarshal([]byte(companyJSON), &data)
		if err != nil {
			fmt.Println("Failed to unmarshall json", err)
		}
		ch <- struct{}{}
	}()
	<-ch
	for key, value := range data {
		fmt.Printf("%s : %v \n", key, value)
	}
}
