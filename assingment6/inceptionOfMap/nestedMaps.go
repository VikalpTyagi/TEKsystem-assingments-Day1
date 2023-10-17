package main

import "fmt"

func main() {
	studentGrades := map[string]map[string]any{
		"Vishal": map[string]any{"age": 27, "Grade": 9.3, "City": "Patna"},
		"pooja":  map[string]any{"age": 27, "Grade": 5.3, "City": "Pune"},
		"Sahil":  map[string]any{"age": 23, "Grade": 9.3, "City": "Mumbai"},
		"Vikalp": map[string]any{"age": 24, "Grade": 9.3, "City": "Kota"},
	}
	for k, v := range studentGrades {
		fmt.Println(k, " : ")
		for k2,v2 := range v{
			fmt.Printf("\t\t%q : %v \n",k2, v2)
		}
	}
}
