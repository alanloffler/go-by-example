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
}
