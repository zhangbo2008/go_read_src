package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 20; i++ {
		n := i

		fmt.Println(n, "===", uint32(-n)%uint32(n))
	}
}

//go:noescape
func Sum(x, y int) int
