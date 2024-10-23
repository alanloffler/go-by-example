// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records

package main

import "fmt"

type person struct {
	name string
	age  int
}

// Constructs a new person struct
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	// Create a new struct
	fmt.Println(person{"Bob", 20})

	// Name fields when initializing struct
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields are zero-valued
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// Encapsulates new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
}
