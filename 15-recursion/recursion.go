package main

import "fmt"

// Factorial recursive function
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// Factorial of 4, calls itself until fact(0)
	fmt.Println(fact(7))

	// Anonymous func must be stored in var
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
