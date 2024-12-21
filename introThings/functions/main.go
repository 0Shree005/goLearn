package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in GO")

	addRes := adder(6, 9)
	fmt.Printf("Addres: %v\n", addRes)

	returnedString, MParamAdd := multiParamAdder(0, 6, 2, 1, 6, 0)
	fmt.Printf("MParamAdd: %v\n", MParamAdd)
	fmt.Printf("returnedString: %v\n", returnedString)
}

func adder(x int, y int) int {
	return x * y
}

func multiParamAdder(values ...int) (string, int) {
	total := 0
	for _, val := range values {
		total += val
	}

	return "You can also return STRINGS", total
}
