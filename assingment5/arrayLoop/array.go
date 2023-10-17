package main

import "fmt"

func main() {
	array := [...]int{10, 02, 34, 48, 95}

	for i := 0; i < len(array)-1; i++ {
		fmt.Printf("Position %d element of array: %d\n", i, array[i])
	}
}
