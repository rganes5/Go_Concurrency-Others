package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name string `json:"name"`
	Year string `json:"year"`
}

func main() {
	company := &Company{
		Name: "GoLinuxCloud",
		Year: "2019",
	}

	//Encoding company data using Marshall function
	jsondata, err := json.Marshal(company)
	if err != nil {
		fmt.Println(err)
	}
	// jsonData variable contains a data in form og bytes.
	// we have handled error in cas if encountered during execution
	fmt.Printf("Bytes : %v \n", jsondata) // will show the data inform of bytes

	// we need to convert the bytes of jsondata into json string by typecast it to a string function
	jsonOutput := string(jsondata)
	fmt.Printf("JSONDATA: %v\n", jsonOutput)
}
