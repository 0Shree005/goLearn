package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "userInput in GO"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for our Pizza: ")

	// comma ok OR error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks your rating, ", input)
	fmt.Printf("Type of the rating is: %T \n", input)

}
