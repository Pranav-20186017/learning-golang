package main

import (
	"fmt"
	"math"
)


func temp() {
	fmt.Printf("Hello from Golang\n")
}



func main() {
	fmt.Printf("hello, world\n")
	c := math.Sqrt(25)
	fmt.Printf("%.1f\n", c)
	temp()
}