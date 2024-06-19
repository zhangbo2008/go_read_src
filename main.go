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
	Overflow := float32(1.7)
	fmt.Println("Overflow:", unsafe.Sizeof(Overflow))
	fmt.Println(2 << 1)
	fmt.Println(4 << 1)
	fmt.Println(8 << 1)
	fmt.Println(16 << 1)
	fmt.Println(32 << 1)
	fmt.Println(32 << (^uint(0) >> 63))
}

//go:noescape
func Sum(x, y int) int
