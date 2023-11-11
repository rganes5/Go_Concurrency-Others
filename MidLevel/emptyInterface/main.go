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
	//Type switch
	switch value := input.(type) {
	case int:
		fmt.Println("This value is of type int", value)
	case string:
		fmt.Println("This value is of type string", value)
	default:
		fmt.Println("Out of earth")
	}
}

// Generics
type custom interface {
	int | string | float64
}

func generics[c custom](input c) {
	fmt.Printf("The value  %v is of type %T \n", input, input)
}

// Aleas of interface *any*
func anyType(input any) {
	fmt.Printf("The value  %v is of type %T \n", input, input)
}

//Type assertion

// You can edit this code!
// Click here and start typing.
/*
func main() {
	var interfaceVal interface{}
	interfaceVal = "This is a string"

	checkTypeString := interfaceVal.(string)
	fmt.Println(checkTypeString)

	checkTypeInt, ok := interfaceVal.(int)
	if ok {
		fmt.Println("Success")
		fmt.Println(checkTypeInt)

	} else {
		fmt.Printf("type is %T", checkTypeInt)

	}

}
*/
