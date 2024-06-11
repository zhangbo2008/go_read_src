package main

import "time"

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	println(b)
	time.Sleep(1 * time.Microsecond)
	time.Sleep(1 * time.Microsecond)
	time.Sleep(1 * time.Microsecond)
	println(a)
}

func main() {
	go f()

	var ffffffffffffffff = 1.234234e12
	println(ffffffffffffffff)

	g()
	time.Sleep(1 * time.Microsecond)
}
