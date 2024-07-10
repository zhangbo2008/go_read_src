package main

func main() {

	c := []byte("daf11")
	a := c[:3]
	b := c[:4]

	println(&a[0]) // 0x17a249
	println(&b[0]) //0x17a249
	println(&a)
	println(&b)

	x := []string{"212", "af"}
	m := make(map[string]int)
	for _, s := range x {
		if c, ok := m[s]; c > -2 {
			print(ok)
			println(m[s])
			m[s] = c - 1
		}
	}
	var b1 interface{}
	b1 = 3
	var b2 = b1.(int)
	print(b2 + 3)
	print(1)

}
