package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@mail.com"
	fmt.Println("Email of the user is:", u.Email)
}

func main() {
	fmt.Println("Structs in GO")

	testUser := User{"User1", "user@mail.com", true, 69}
	fmt.Println(testUser)
	fmt.Printf("User's details are: %v\n", testUser)
	fmt.Printf("Name: %v\nEmail:%v\nLoggedIn:%v\nAge:%v\n", testUser.Name, testUser.Email, testUser.Status, testUser.Age)
	testUser.GetStatus()
	testUser.NewMail()                                          // THIS here modifies the mail but it doesn't affect the testUser.Email, pass by value
	fmt.Println("Email ealier initialised is:", testUser.Email) // IT still has the user@mail.com it wasn't overwritten by the method NewMail

}
