package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create empty map
	m := make(map[string]int)

	// Set values
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// Get value by key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Zero value for non-existent key
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Lenght of keys/values
	fmt.Println("len:", len(m))

	// Remove key/value
	delete(m, "k2")
	fmt.Println("map:", m)

	// Remove all keys/values
	clear(m)
	fmt.Println("m:", m)

	// Second optional value
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize in same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)

	// Maps utility package
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
