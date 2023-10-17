package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30,88}
	sum := calSum(intSlice)
	fmt.Println("Sum even of elements is: ", sum)

}
func calSum(slice []int) int {
	sum:= 0
	for _, num := range slice {
		if num%2==0 {
			sum = sum + num
		}
	}
	return sum
}
