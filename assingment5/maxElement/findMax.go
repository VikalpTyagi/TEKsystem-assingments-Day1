package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30}
	fmt.Println("Largest value in slice is: ", findMax(intSlice))
}

func findMax(slice []int) int {
	var max int
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[i] > slice[j+1] {
				max = slice[i]
			}
		}
	}
	return max
}
