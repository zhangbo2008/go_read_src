package main

func main() {

	//...
	// 不知道为啥拷出来
	var i = 3333
	println(i)
	var a1 = &i
	println(a1)
	println(*a1)

}

//go:noescape
func Sum(x, y int) int
func Sum2(x, y int) int
