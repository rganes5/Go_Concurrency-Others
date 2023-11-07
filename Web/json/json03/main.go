package main

import (
	"encoding/json"
	"fmt"
)

var companyJSON = `{
	"name" : "GOLinuxCloud",
        "years_of_service" : "5",
        "nature_of_company" : "Online Academy",
        "no_of_staff" : 10
}`

type company struct {
	Name              string
	Years_of_service  string
	Nature_of_company string
	No_of_staff       int32
}

func main() {
	var companyData company
	err := json.Unmarshal([]byte(companyJSON), &companyData)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("Name:", companyData.Name)
	fmt.Println("Years of Service:", companyData.Years_of_service)
	fmt.Println("Nature of Company:", companyData.Nature_of_company)
	fmt.Println("No. of Staff:", companyData.No_of_staff)
}
