package main

import "fmt"

func main() {
	var arr[2] string
	arrDemo := [...] string{"Demo", "Example", "Hello"}
	arr[0] = "Hello"
	arr[1] = "Demo"

	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)
	fmt.Printf("%q \n", arr)

	fmt.Printf("%q \n", arrDemo)
}