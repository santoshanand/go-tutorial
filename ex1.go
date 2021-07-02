package main 

import "fmt"

type Artist struct {
	Name string
	Genre string
	Songs int
}

func NewRelease(artist Artist) int {
	artist.Songs ++
	return artist.Songs
}

func main() {
	res := Artist{Name: "Santosh", Genre: "Electro", Songs: 12}

	fmt.Printf("%s Name %s Genre %d Songs \n", res.Name, res.Genre, NewRelease(res))
	fmt.Printf("%s Name %s Genre %d Songs \n", res.Name, res.Genre, res.Songs)

}