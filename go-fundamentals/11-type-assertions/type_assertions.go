package main

import (
	"fmt"
)

type MyInterface interface {
	Method()
}

type MyStruct struct{}

func (m *MyStruct) Method() {
	fmt.Println("Method called")
}

func main() {
	var i MyInterface // i is an interface holding a nil value

	// Here, we attempt to assert i to a concrete type
	if _, ok := i.(*MyStruct); ok {
		fmt.Println("i holds a MyStruct")
	} else {
		fmt.Println("i does NOT hold a MyStruct")
	}

	// Assigning a nil pointer of MyStruct to the interface
	i = (*MyStruct)(nil)

	// Now we perform a type assertion again
	if _, ok := i.(*MyStruct); ok {
		fmt.Println("i holds a MyStruct")
	} else {
		fmt.Println("i does NOT hold a MyStruct")
	}

	// Attempting to call a method on the nil pointer
	if m, ok := i.(*MyStruct); ok {
		m.Method() // This will cause a panic if m is nil
	} else {
		fmt.Println("Cannot call Method, i is not a MyStruct")
	}
}
