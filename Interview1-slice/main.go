package main

import "fmt"

// array
func modifyArray(array [5]int) {
	array[2] = 50
	array[3] = 100
}

// slice
func modifySlice(slice []int) {
	slice[2] = 50
	slice[3] = 100
}

// main
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before modification", array)
	modifyArray(array)
	fmt.Println("After modification", array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before modification", slice)
	modifySlice(slice)
	fmt.Println("After modification", slice)

}
