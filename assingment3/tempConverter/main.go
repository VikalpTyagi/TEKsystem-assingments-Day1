package main

import (
	"fmt"
	"tempAssing/temprature"
)

func main() {
	fmt.Println("Enter the Temperature:")
	var input int
	_,err:= fmt.Scanln(&input)
	if err != nil{
		fmt.Println("Facing Error :",err)
		return
	}
	temp :=temprature.DegreeConverter(input)
	fmt.Println("Temperature in Celcius = ",temp)

}