package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}

	fmt.Println(comparator(slice1, slice2))

}

func comparator(s1, s2 []int) string {
	if len(s1)!=len(s2) {
		return "Both the slices are different"
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return "Both the slices are different"
		}
	}
	return "Both the slices are same!!"
}
