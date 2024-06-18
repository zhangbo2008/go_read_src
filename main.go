package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 10
	y := 20
	sum := Sum(x, y)
	fmt.Println("Sum:", sum)
	var a = 100
	fmt.Println("a:", unsafe.Sizeof(a))
}

//go:noescape
func Sum(x, y int) int
