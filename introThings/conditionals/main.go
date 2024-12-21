package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditionals in GO")

	loginCount := 1

	var result string

	if loginCount < 10 {
		result = "if"
	} else if loginCount > 10 { // the else and else if statements HAVE to be on the same line on which the IF closes
		result = "else if"
	} else {
		result = "else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number if even")
	} else {
		fmt.Println("Number if odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

	// if err != nil {
	// }

}
