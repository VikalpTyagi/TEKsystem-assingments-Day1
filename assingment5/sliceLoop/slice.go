package main

import "fmt"

func main() {
	slice := []string{"mangoes", "apples", "pineapples", "watermelon"}
	for i,fruitName := range slice {
		fmt.Printf("%d Name of fruits I like: %s \n",i,fruitName)
	}
}
