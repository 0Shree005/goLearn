package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Mytime")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) // dd/mm/yyyy format

	dateCreation := time.Date(2069, time.March, 04, 13, 21, 0, 0, time.UTC)

	fmt.Println(dateCreation)
	fmt.Println(dateCreation.Format("02-01-2006 15:04:05 Monday"))

}
