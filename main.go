package main

import (
	_ "embed"
	"fmt"
	"math"
	"unsafe"
)

//go:embed hello.txt
var b []byte

func main() {
	fmt.Println(b)
	var a = math.NaN()
	print(a != a)

	var firstStoreInProgress byte
	print(unsafe.Pointer(&firstStoreInProgress))
	for i := 1; i <= 20; i++ {
		n := i

		fmt.Println(n, "===", uint32(-n)%uint32(n))
	}
	d := make(chan struct{})
	print(d)
}

//go:noescape
func Sum(x, y int) int
func Sum2(x, y int) int
