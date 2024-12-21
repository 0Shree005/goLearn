package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers in GO")

	// var ptr *int
	// fmt.Println("The default value of a pointer is: ", ptr) // <nil>
	// fmt.Println("derefrencing a nil pointer", *ptr) // panic: runtime error: invalid memory address or nil pointer dereference

	intPtr := 30

	var ptr = &intPtr

	fmt.Println("Address of intPtr:", ptr)
	fmt.Println("Value of intPtr:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Modified value of ptr is:", intPtr)
}
