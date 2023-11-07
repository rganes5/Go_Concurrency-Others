package main

import (
	"encoding/json"
	"fmt"
)

// sample json object to be parsed
var companyJSON = `[{"name" : "GOLinuxCloud","year" : "2000"},
{"name" : "Google","year" : "1998"}]`

type Company struct {
	Name string
	Year string
}

func main() {
	var company []Company
	err := json.Unmarshal([]byte(companyJSON), &company)
	if err != nil {
		fmt.Println("err", err)
	}
	for _, val := range company {
		fmt.Printf("Company : %+v \n", val)

	}
	// Print array of values in the struct
	//fmt.Printf("Company : %+v \n", company)
}
