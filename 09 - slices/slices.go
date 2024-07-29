package main

import (
	"fmt"
)

func main() {
	// Uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Empty slice with non-zero length
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Set and get
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Length
	fmt.Println("len:", len(s))

	// Append values to slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice operations
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and initialize
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Composed multi-dimensional slides
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
