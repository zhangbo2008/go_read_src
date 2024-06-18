package main

import "fmt"

func main() {
	x := 10
	y := 20
	sum := Sum(x, y)
	fmt.Println("Sum:", sum)
}

//go:noescape
func Sum(x, y int) int
