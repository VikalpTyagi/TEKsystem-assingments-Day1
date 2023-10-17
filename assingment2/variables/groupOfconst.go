package main

import "fmt"

func main() {
	const (
		name string = "Vikalp"
		age  int    = 24
	)
	fmt.Printf("name: %q \nage: %d",name, age)

	//Line below gives error because you can not reassign value to const variables
	// name = "Ajay"

}