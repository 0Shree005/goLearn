package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Routines")

	go greeter("routine 1")
	go greeter("routine 2")

	time.Sleep(3 * time.Second)
	fmt.Println("Main functions ends!")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %v, s: %s\n", i, s)
		time.Sleep(500 * time.Millisecond)
	}
}
