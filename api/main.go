package main

import (
	"fmt"

	"api/hii"
)

// Embedding of struct

/*
type Student struct {
	name  string
	age   int
	place string
	Class
}

type Class struct {
	batch string
}

func main() {
	student := Student{
		name:  "Ganesh",
		age:   25,
		place: "idukky",
		Class: Class{
			batch: "BCA",
		},
	}

	fmt.Println(student.batch)
}
*/

//Interface
/*
type database interface {
	findUser() string
}

func Usecase(d database) {

	usernName := d.findUser()
	fmt.Println(usernName)
}

type postgresdb struct {
}

func (c *postgresdb) findUser() string {
	return "postgres user"
}

type mongodb struct {
}

func (c *mongodb) findUser() string {
	return "mongo user"
}

func main() {

	Usecase(&postgresdb{})
	Usecase(&mongodb{})
}
*/

// anonymous fn

/*
func main() {

	takingAfunction(func() {
		fmt.Println("hiii")
	})
}

// higherorder fn
func takingAfunction(f func()) {
	f()
}

func hii() {
	fmt.Println("hiii")
} */

// encapsulation

func main() {

	h := hii.GetSecret("key")
	h.Name = "hii"

	s := hii.Secret{
		Name: "Nikhil",
	}
	fmt.Println(s)

	fmt.Println(h.GetPassword("hii"))
}
