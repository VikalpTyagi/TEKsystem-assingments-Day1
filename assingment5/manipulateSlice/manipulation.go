package main

import "fmt"

func main() {
	intSlice := []int{}
	intSlice = append(intSlice, 5)
	fmt.Println("5 appended",intSlice)
	intSlice = append(intSlice, 10)
	fmt.Println("10 appended",intSlice)
	intSlice = append(intSlice, 15)
	fmt.Println("15 appended",intSlice)
	intSlice = append(intSlice, 20)
	fmt.Println("20 appended",intSlice)
	intSlice = append(intSlice, 25)
	fmt.Println("25 appended",intSlice)

	fmt.Println("Slice capacity =", cap(intSlice))
	fmt.Println("Slice length =", len(intSlice))

	sum := calSumOfSlice(intSlice)
	fmt.Println("Sum of all the integer in slice is: ", sum)

	intSlice = deleteElement(intSlice,2)
	fmt.Println("After removing 15 \n",intSlice)
}

func deleteElement(slice []int, i int) []int{
	return append(slice[:i],slice[i+1:]...)
}

func calSumOfSlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum = sum + num
	}
	return sum
}
