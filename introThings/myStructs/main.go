package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in GO")

	testUser := User{"User1", "user@mail.com", true, 69}
	fmt.Println(testUser)
	fmt.Printf("User's details are: %v\n", testUser)
	fmt.Printf("Name: %v\nEmail:%v\nLoggedIn:%v\nAge:%v\n", testUser.Name, testUser.Email, testUser.Status, testUser.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
