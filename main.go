package main

import (
	"fmt"
)

type asciiSet [8]uint32

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1, 20: 1}

func main() {
	for index, value := range asciiSpace {
		if value != 0 {
			fmt.Println(index, value)
		}
	}
}
