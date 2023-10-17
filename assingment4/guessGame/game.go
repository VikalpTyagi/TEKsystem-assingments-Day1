package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(10)
	tries:= 0
	fmt.Println("target: ",target)
	fmt.Println("*******************************START**************************************")

	for{
		tries ++
		
		var guess int
		fmt.Printf("Guess the number from 1 to 10\ttries: %d\n",tries)
		_,err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Error: ",err)
			return
		}
		if guess == target {
			fmt.Printf("-------You guessed correctly!!-------\nNumber: %d\n\n",target)
			if tries == 1{
				fmt.Println("but you cheated!!")
			}
	fmt.Println("*******************************FINISHED**************************************")

			return
		}else if guess > target {
			fmt.Println("-----Too high!, Try again!")
		}else if guess < target {
			fmt.Println("-----Too low!, Try again!")
		}

	}
}