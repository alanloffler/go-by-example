package main

import "fmt"

// Function returns double int
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Returns double int
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Returns a subset (second value in this case)
	_, c := vals()
	fmt.Println(c)
}
