// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	animal := Animal{
		Area: "Land",
		Dog: Dog{
			Name: "Dog",
			Legs: 4,
		},
	}
	animal.Greet()

}

type Animal struct {
	Dog
	Area string
}

type Dog struct {
	Name string
	Legs int
}

func (a *Animal) Greet() {
	fmt.Printf("THe name of the animal is %s and it lives in %s and it has %d legs", a.Name, a.Area, a.Legs)
}
