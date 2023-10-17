package main

import (
	"fmt"
)

func main() {
	intSlice := []int{1, 1, 3, 4, 5, 3, 5, 56, 3}
	
	unique := dedepulicate(intSlice)
	fmt.Println("Dedepulicated slice: ", unique)
}

func dedepulicate(slice []int) []int {
	list := make([]int,0, len(slice))
	for i := 0; i < len(slice); i++ {
		flag:=false
		for j := 0; j < len(list); j++ {
			if slice[i] == list[j] {
				flag = true
				break
			}
		}
		if !flag{
			list = append(list, slice[i])
		}
	}
	return list
}
