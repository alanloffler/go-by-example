package main

import "fmt"

// Closure: returns anonymous function that closes over var i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// First call
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Second call
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())

	// First call again
	fmt.Println(nextInt())
}
