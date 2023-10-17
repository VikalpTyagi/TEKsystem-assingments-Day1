package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30}
	sum, avg := calSumAndAverage(intSlice)
	fmt.Printf("Sum of elements is: %d\nAverage is: %d", sum, avg)

}
func calSumAndAverage(slice []int) (int, int) {
	sum, avg := 0, 0
	for _, num := range slice {
		sum = sum + num
	}
	avg = sum / len(slice)
	return sum, avg
}
