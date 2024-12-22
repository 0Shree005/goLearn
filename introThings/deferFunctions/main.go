package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer Functions in GO")

	// defer fmt.Println("World")
	// fmt.Println("Hello")

	// any thing prefixed by the ( defer ) keyword is STACKED just before the func ends
	// WORLD will be stacked HERE, and will be printed AFTER hello in the terminal as the control flow here is 1.HELLO, 2. WORLD

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

	// now for this, as they are STACKED it follows the LIFO thing, so the order will be, World, One, Two
	// so now when printing it to the it'll be, 1. Hello, 2. Two, 3. One, 4. World
}

func myDefer() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// hello, World, One, Two, newLine, 0, 1, 2, 3, 4
// hello, newline, 4 3 2 1 0, two, one, world
