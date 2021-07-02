package main

import "fmt"

func main() {
	tempSlice := []int{1,3,5,6,7,8, 10}
	fmt.Println(tempSlice)
	fmt.Println(tempSlice[1:3])
	fmt.Println(tempSlice[:5])
	fmt.Println(tempSlice[2:])
}