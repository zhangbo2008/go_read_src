package main

func main() {
	x := 0xf1                     //   0....011110001
	println(int32(x) << 31 >> 31) // 11111111111 32个

}
