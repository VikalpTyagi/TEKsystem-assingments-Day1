package main

import "fmt"

func main() {
	intSlice := []int{10, 15, 20, 25, 30}
	var num int
	fmt.Println("Enter the number to find :")
	_,err :=fmt.Scanln(&num)
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println(findNumber(intSlice,num))
}

func findNumber(slice []int,num int) string{
for _,ele:= range slice{
	if ele == num {
		return "Element found successfully!!"
	}
}
return "Element doesn't exisits in list"

}
