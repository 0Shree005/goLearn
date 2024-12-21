package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in GO")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and the day is %v\n", index+1, day)
	// }

	// goto
	initialVal := 0

	for initialVal < 10 {

		if initialVal == 2 {
			goto web
		}

		if initialVal == 5 {
			initialVal++
			continue
		}
		fmt.Println("Value is: ", initialVal)
		initialVal++
	}

web:
	fmt.Println("github.com/0Shree005")

}
