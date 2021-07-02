// This logic has O(n) complexcity

package main

import "fmt"

func search(arr []int, number int)int {
	for i, v := range arr {
		if number == v {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{121,12, 232,45,54,675,66}
	number := 123

	fmt.Println(search(arr, number))
}