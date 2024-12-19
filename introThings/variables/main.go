package main

import "fmt"

// you can't do the following
// numbersOUTSIDE := 2340324
// but can do this instead, which doing so creates a PUBLIC variable
var PublicVariable string = "Public"

// or, which creates a constant
const RandomVar string = "awldiawd"

func main() {
	var username string = "bogo"
	fmt.Printf("Var is of type: %T \n", username)

	var isGo bool = true
	fmt.Printf("Var is of type: %T \n", isGo)

	var smallVal uint8 = 255
	fmt.Printf("Var is of type: %T \n", smallVal)

	// a check for what the default values are for these types

	var defaultCHKstring string
	fmt.Println(defaultCHKstring)
	fmt.Printf("Var is of type: %T \n", defaultCHKstring)

	var defaultCHKint int
	fmt.Println(defaultCHKint)
	fmt.Printf("Var is of type: %T \n", defaultCHKint)

	var defaultCHKfloat32 float32
	fmt.Println(defaultCHKfloat32)
	fmt.Printf("Var is of type: %T \n", defaultCHKfloat32)

	var defaultCHKfloat64 float64
	fmt.Println(defaultCHKfloat64)
	fmt.Printf("Var is of type: %T \n", defaultCHKfloat64)

	// you can do this tho
	numbers := 3.0
	fmt.Printf("Var is of type: %T \n", numbers)

}
