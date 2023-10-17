package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := concatenateSlices(slice1, slice2)
	fmt.Println("Combined slice: ", slice3)

}

func concatenateSlices(s1, s2 []int) []int {
	return append(s1, s2...)
}
