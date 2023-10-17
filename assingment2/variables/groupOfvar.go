package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		height float32
	)

	name = "Vikalp"
	age = 24
	height = 6.1

	fmt.Printf("Name : %q \nAge: %d\nHeight: %g", name, age, height)
}
