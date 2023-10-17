package main

import (
	"calAssingment/calculator"
	"fmt"
)

func main() {
	sum := calculator.Sum(2,2)
	fmt.Println("Sum is: ",sum)

	difference := calculator.Differentiator(5,2)
	fmt.Println("Difference is: ",difference)

	square := calculator.Square(4)
	fmt.Println("Square is: ",square)

	q, r := calculator.Divider(2, 8)
	fmt.Printf("Quotient: %d\nRemainder: %d\n",q,r)

	
	fmt.Println("input 1st value for multiplication:")
	var input1 int
	_,err := fmt.Scanln(&input1)
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("input 2nd value for multiplication:")
	var input2 int
	_,err = fmt.Scanln(&input2)
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	multiple := calculator.Multiplication(input1,input2)
	fmt.Println("Multiplication is: ",multiple)
}
