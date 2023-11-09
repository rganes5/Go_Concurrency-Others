package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	rec1 := Rectangle{}
	rec1.Length = 3.4
	rec1.Width = 3.4
	fmt.Println(rec1)
	fmt.Println("rectangle 1:")
	fmt.Printf("Length: %v & Width: %v \n", rec1.Length, rec1.Width)

	rec2 := Rectangle{4.5, 5.6}
	fmt.Println(rec2)
	fmt.Println("Area of rect 2 is ", rec2.Area())

	rec3 := &Rectangle{
		Length: 4.5,
		Width:  9.5,
	}
	fmt.Println(rec3)
	fmt.Println("Perimeter of rect 3 is ", rec3.Perimeter())

	bob := Employee{
		ID:     1,
		Salary: 10000,
		Person: Person{
			Name: "Bob",
			Age:  25,
		},
	}

	fmt.Println("BOB:", bob)

	//to rep as json
	Ganesh := Person{
		Name: "Ganesh",
		Age:  25,
	}
	//Marshalling
	byteSlice, err := json.Marshal(Ganesh)
	if err != nil {
		fmt.Println("error marshalling", err.Error())
	}
	fmt.Println("JSON", byteSlice)
	fmt.Printf("JSON string %s \n", byteSlice)

}

type Rectangle struct {
	Length float32
	Width  float32
}

// method with val receiver
func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

// method with pointer receiver
func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Width)
}

// Field tags
type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	//
}

type Employee struct {
	ID     uint64
	Person Person
	Salary uint64
}
