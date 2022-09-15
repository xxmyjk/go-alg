package main

import (
	matrix "github.com/xxmyjk/go-alg/pkg/matrix"
)

func main() {
	el := []int{
		1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21,
		50, 23, 24, 25, 26, 27, 28,
	}

	m := matrix.New(4, 7, el...)
	// m.Print()
	m.CircleLikePrint()
}
