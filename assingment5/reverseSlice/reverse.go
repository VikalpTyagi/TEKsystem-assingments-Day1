package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30}
	fmt.Println("Reverse slice is: ", reverse(intSlice))
}

func reverse(slice []int) []int {
	// temp := 0
	leng := len(slice)
	for i := 0; i < leng/2; i++ {
		// temp = slice[leng-i-1]
		// slice[leng-i-1] = slice[i]
		// slice[i] = temp
		slice[leng-i-1], slice[i] = slice[i], slice[leng-i-1]
	}
	return slice
}
