package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _,v := range a {
		sum += v
	}

	c <- sum
}

func main() {
	value := []int {1,4,6,7,-5,33}
	c := make(chan int)

	go sum(value[:len(value)/2], c)
	go sum(value[len(value)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}