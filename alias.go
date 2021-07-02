package main

import "fmt"
import "strings"

type StringDemo string


func (s StringDemo) Uppercase() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println(StringDemo("demo").Uppercase())
}