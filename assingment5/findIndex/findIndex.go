package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30, 88}
	number := 18 //Number of which we have to finde index value of
	index := findIndex(intSlice,number)
	if index < 0 {
		fmt.Printf("Can't find the number: %d\nIn list: %v",number, intSlice)
		return
	}
	fmt.Printf("The number: %d\nIndex value is: %d", number, index)
}

func findIndex(slice []int, num int) int {
	for i, ele := range slice {
		if num == ele {
			return i
		}
	}
	return -1
}
