package main

import "fmt"

func main() {
	studentGrades := map[string]float32{
		"Vishal":     8.9,
		"pooja":     8.9,
		"Sahil":      9.0,
		"Vikalp":     9.9,
	}
	for k,v := range studentGrades{
		fmt.Println(k," : ",v)
	}
	delete(studentGrades,"pooja")
	fmt.Printf("\n\nAfter deletion of pooja: \n")
	for k,v := range studentGrades{
		fmt.Println(k," : ",v)
	}
}