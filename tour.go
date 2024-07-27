package main

import (
	"fmt"
	"math"
)

// Constants
const s string = "constant"

func main() {
	// hello world
	fmt.Println("---------- Hello World ----------")
	fmt.Println("Hello, World!")

	// Values
	fmt.Println("---------- Values ----------")
	fmt.Println("Go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variables
	fmt.Println("---------- Variables ----------")
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println("---------- Constants ----------")
	fmt.Println(s)

	const n = 500000000
	const d2 = 3e20 / n
	fmt.Println(d2)

	fmt.Println(int64(d2))

	fmt.Println(math.Sin(n))
}
