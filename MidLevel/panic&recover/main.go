package main

import "fmt"

func main() {
	Demo(100, 0)
	Stupid()
}

func Demo(a, b int32) {
	defer handlePanic()
	if b == 0 {
		panic("Division by zero")
	} else {
		c := a / b
		fmt.Println(c)
	}
}

func Stupid() {
	defer handlePanic()
	fmt.Println("Something bad about to happen")
	panic("Ending the program")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("RECOVER", r)
	}
}
