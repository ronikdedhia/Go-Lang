package main

import (
    "fmt"
)

func addMulti(x, y int) (int, int) {
    return (x + y), (x * y)
}

func main() {
    var i int
 	var j int
    fmt.Print("Type a number: ")
    fmt.Scan(&i)
 	fmt.Print("Type a number: ")
  	fmt.Scan(&j)
    a, b := addMulti(i, j)
    fmt.Printf("Addition is %d & Multiplication is %d\n", a, b)
}   