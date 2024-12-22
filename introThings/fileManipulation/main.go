package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File manipulations in GO")
	content := "Data to be written in the fileName"

	var fileName string = "./testFile.txt"

	file, err := os.Create(fileName)

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile(fileName)
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	fmt.Printf("File data: %v", (data))

	checkNilError(err)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
