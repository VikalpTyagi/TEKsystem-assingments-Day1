package main

import "fmt"

func main() {
	num1, num2 := 844, 999
	fmt.Println("For number: ",num1)
	fmt.Println(chekerEvenOdd(num1))
	fmt.Println("\nFor number: ",num2)
	fmt.Println(chekerEvenOdd(num2))

}

func chekerEvenOdd(num int) string {
	if num%2 == 0 {
		return "Number is even!"
	}
	return "Number is odd!"
}
