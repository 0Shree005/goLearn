package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in GO")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(7)
	fmt.Println("Dice number is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can move 1 place")
	case 2:
		fmt.Println("Dice value is 2 and you can move 2 places")
	case 3:
		fmt.Println("Dice value is 3 and you can move 3 places")
	case 4:
		fmt.Println("Dice value is 4 and you can move 4 places")
	case 5:
		fmt.Println("Dice value is 5 and you can move 5 places")
	case 6:
		fmt.Println("Dice value is 6 you can OPEN and roll again")
	default:
		fmt.Println("Invalid")
	}
}
