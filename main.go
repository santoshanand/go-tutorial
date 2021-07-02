package main 

import "fmt"
func main() {
  mapDemo()
  fmt.Println("Go working")
}


func mapDemo() {
  mapVar := make(map[string]string)
  mapVar["demo1"] = "Demo one"
  mapVar["demo2"] = "Demo two"
  fmt.Println(mapVar)
}