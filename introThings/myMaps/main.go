package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in GO")

	langs := make(map[string]string) // key value pair things

	langs["GO"] = "GoLang"
	langs["JS"] = "Javascript"
	langs["PY"] = "Python"
	langs["RB"] = "Ruby"

	fmt.Println("List of all the langs:", langs)
	fmt.Println("JS is an abv of:", langs["JS"]) // can access the value using the key

	delete(langs, "PY") // will remove Python from the map using the key
	fmt.Println("List of all the langs:", langs)

	// looping through the map using the key and val, kinda like for each loop
	for key, value := range langs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
