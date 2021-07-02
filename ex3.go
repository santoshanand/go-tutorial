package main

import "fmt"

type User struct {
	Name string
	Age int
	Phone string
}


type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
)
func main() {

	userNew := new(User)
	userAnd := &User{}

	user := User{Name: "Santosh", Phone: "343434", Age: 2}

	fmt.Println(user)

	fmt.Println(p)
	fmt.Println(q)

	fmt.Println(*userNew == *userAnd)
}