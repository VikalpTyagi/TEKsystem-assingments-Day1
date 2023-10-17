package main

import "fmt"

func main() {
	fruitsNames := map[string]int{
		"Mango":  5,
		"Apples": 12,
		"Guava":  8,
		"banana": 6,
	}
	for k, v := range fruitsNames {
		fmt.Println(k, ",Quantity: ", v)
	}

}
