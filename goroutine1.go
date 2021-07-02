package main 

import (
	"fmt"
	"time"
)

func sayHello(v string) {
	for i:= 0; i<5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(v)
	}
}

func main() {
	go sayHello("Hi")
	sayHello("Welcome")
}