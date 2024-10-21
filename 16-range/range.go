package main

import "fmt"

func main() {
	// Range iterates over elements of an array
	nums := []int{2, 3, 4}
	sum := 0

	// Without index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// With index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on map, iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Iterates only over keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Iterates only over values
	for _, v := range kvs {
		fmt.Println("value:", v)
	}

	// Range on strings iterates over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
