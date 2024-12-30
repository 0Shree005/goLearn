package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	fmt.Println("Go Routines")

	// go greeter("routine 1")
	// go greeter("routine 2")

	time.Sleep(3 * time.Second)

	websitelist := []string{
		"https://go.dev",
		"https://github.com",
		"https://google.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
	fmt.Println("Main functions ends!")
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error in endpoints")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

//
// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("i: %v, s: %s\n", i, s)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
