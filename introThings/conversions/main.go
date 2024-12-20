package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 & 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your rating: ")

	userInput, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your rating of: ", userInput)
	fmt.Printf("The type of the rating is: %T \n", userInput)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

	fmt.Printf("The type of the rating is: %T \n", numRating)
}
