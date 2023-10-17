package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30}
	fmt.Println("New slice is: ",doubleValueElemets(intSlice))
}

func doubleValueElemets(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] += slice[i] 
	}
	return slice
}
