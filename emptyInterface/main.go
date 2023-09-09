package main

import "fmt"

func main() {
	Hello(2.5)
	generics(50)
	generics("HELLO")
	generics(4.3)
	anyType(4)
	anyType("string")

}

func Hello(input interface{}) {

	switch value := input.(type) {
	case int:
		fmt.Println("This value is of type int", value)
	case string:
		fmt.Println("This value is of type string", value)
	default:
		fmt.Println("Out of earth")
	}
}

type custom interface {
	int | string | float64
}

func generics[c custom](input c) {
	fmt.Printf("The value  %v is of type %T \n", input, input)
}

func anyType(input any) {
	fmt.Printf("The value  %v is of type %T \n", input, input)

}
