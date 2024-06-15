package main

import (
	"math"
)

const zero = 121.11

func main() {

	x := -float64(zero)
	b := math.Float64bits(x)
	print(b)
}
