package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	CircumF() float32
}

type Circle struct {
	radius float32
}

//Circle methods

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) CircumF() float32 {
	return 2 * c.radius * math.Pi
}

type Square struct {
	Length float32
}

//Square methods

func (s Square) Area() float32 {
	return s.Length * s.Length
}

func (s Square) CircumF() float32 {
	return s.Length * 4
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area of %T is :%0.2f \n", s, s.Area())
	fmt.Printf("Circumference of %T is :%0.2f \n", s, s.CircumF())
}

func main() {
	shape := []Shape{
		Square{Length: 2.4},
		Circle{radius: 3.4},
		Square{Length: 4.5},
		Circle{radius: 5.6},
	}

	for _, v := range shape {
		printShapeInfo(v)
		fmt.Println("----")
	}
	/*
		sq := Square{
			4.0,
		}

		ci := Circle{
			3.2,
		}
	*/
}
