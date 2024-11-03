// Interfaces are named collections of method signatures

package main

import (
	"fmt"
	"math"
)

// Interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// Rect and circle types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, needs to implement all the methods in the interface
// Rect methods
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If variable has an interface type, then we can call methods defined in his interface
func measure(g geometry) {
	fmt.Println("geometry:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
