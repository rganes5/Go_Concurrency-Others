package main

import "fmt"

func main() {
	rec1 := Rectangle{}
	rec1.Length = 3.4
	rec1.Width = 3.4
	fmt.Println(rec1)

	rec2 := Rectangle{4.5, 5.6}
	fmt.Println(rec2)

	rec3 := Rectangle{
		Length: 4.5,
		Width:  9.5,
	}
	fmt.Println(rec3)

	bob := Employee{
		ID:     1,
		Salary: 10000,
		Person: Person{
			Name: "Bob",
			age:  25,
		},
	}

	fmt.Println("BOB", bob)

}

type Rectangle struct {
	Length float32
	Width  float32
}

type Person struct {
	Name string
	age  uint8
	//
}

type Employee struct {
	ID     uint64
	Person Person
	Salary uint64
}
