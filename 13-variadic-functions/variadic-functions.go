package main

import "fmt"

// Variadic function: any number of trailing arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Two args
	sum(1, 2)

	// Three args
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}

	// Two args by position
	sum(nums[0], nums[3])

	// All args
	sum(nums...)
}
