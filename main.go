package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	var firstStoreInProgress byte
	print(unsafe.Pointer(&firstStoreInProgress))
	for i := 1; i <= 20; i++ {
		n := i

		fmt.Println(n, "===", uint32(-n)%uint32(n))
	}

	var mu sync.Mutex
	var i int

	// 第一次加锁放锁
	mu.Lock()
	//...
	// 不知道为啥拷出来
	m := mu

	i += 1
	m.Unlock()

	// 第二次加锁放锁

	i += 1
	mu.Unlock()

}

//go:noescape
func Sum(x, y int) int
func Sum2(x, y int) int
