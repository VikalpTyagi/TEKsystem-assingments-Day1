package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30, 35, 40, 45,66}
	even, odd := countEvenNOdd(intSlice)
	fmt.Printf("Number of even elements: %d\nNumber of odd element: %d",even,odd)
}

func countEvenNOdd(slice []int) (int, int) {
	countO, countE := 0, 0
	for _, num := range slice {
		if num%2 == 0 {
			countE++
		}else{countO++}
		
	}
	return countE, countO
}
