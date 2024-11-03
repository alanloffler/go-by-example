// Go support methods defined on struct types

package main

import "fmt"

type rect struct {
	width, height int
}

// Area method is a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	var r rect = rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles pointer or value method calls (Pointer avoid copying on method calls or allow method to mutate the struct)
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
