package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d2 = 3e20 / n
	fmt.Println(d2)

	fmt.Println(int64(d2))

	fmt.Println(math.Sin(n))
}
