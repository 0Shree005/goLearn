package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in GO")

	// declaring an array of size 4
	var indexDeclaration [4]string // it initialises 4 values all `empty`

	indexDeclaration[0] = "sflij"
	indexDeclaration[2] = "wasli"
	indexDeclaration[3] = "gijrg"

	fmt.Println("indexDeclaration is:", indexDeclaration)
	fmt.Println("length of indexDeclaration is:", len(indexDeclaration)) // will show 4 as length even tho there are 3 total elements in the array, because it was declared as a size 4 array

	// declaring a predefined array of size 3
	var inlineDeclaration = [3]string{"awd", "sef", "drg"} // it has all the values initialised as well inline with the array's declaration
	fmt.Println("length of inlineDeclaration is: ", len(inlineDeclaration))
}
