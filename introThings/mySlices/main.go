package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4) // allocates memory for a slice of type INT of size 4

	// the index values which were allocated are initialised
	highScores[0] = 321
	highScores[1] = 234
	highScores[2] = 532
	highScores[3] = 983

	highScores = append(highScores, 123, 456, 789) // new elements can be added to a slice EVEN tho it wasn't accounted for initially, when the slice was created

	// fmt.Println(highScores)

	sort.Ints(highScores) // time complexity: O(nlogn)
	// fmt.Println(highScores)
	// fmt.Println("Is sorted: ", sort.IntsAreSorted(highScores))

	// remove an element from slices based on the index
	var langs = []string{"Go", "C", "C++", "Python", "Rust", "Haskell"}
	fmt.Println(langs)
	var index int = 2
	fmt.Println(langs[2])                             // second index has C++
	langs = append(langs[:index], langs[index+1:]...) // this won't print C++ as the slices START from 0 till the index (2), WHICH then start from index+1 = 3, so skipping the second index, and goes till the end, soo the second index value C++ has been REMOVED from the slices, using slice ranges
	fmt.Println(langs[2])                             // second index NOW has python

}
