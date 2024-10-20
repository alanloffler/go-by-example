package main

import (
	"fmt"
)

func main() {
	var a [5]int
	// Prints empty array. Non declared values are set to 0
	fmt.Println("emp:", a)

	// Set and get value
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Length
	fmt.Println("len:", len(a))

	// Declare and initialize 1
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Declare and initialize 2
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Zeroed
	b = [5]int{100, 2: 400, 4: 500}
	fmt.Println("dcl:", b)

	// Composed for multidimensional
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Create and initilize
	var twoD2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", twoD2)
}
