// package main
//
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
//
// type Result struct {
// 	fileName       string
// 	fileChangeBool bool
// }
//
// func main() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	fmt.Printf("Type of ticker is %T\n", ticker)
// 	fmt.Printf("Type of tickerChan is %T\n", ticker.C)
// 	fmt.Println("Process Started!")
// 	fileChangeChan := make(chan Result)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	ch := make(chan bool)
//
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				fmt.Println("Tick")
// 				go printOne(ch, &wg)
// 			case <-Result.C:
// 				fmt.Println("File name is: %s", Result.fileName)
// 				fmt.Println("FileChange was: %v", Result.fileChangeBool)
// 			}
// 		}
// 	}()
//
// 	go printTwo(ch, &wg)
//
// 	wg.Wait()
// 	fmt.Println("Process ends!")
// }
//
// func printOne(ch chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	timeout := 5 * time.Second
//
// 	select {
// 	case <-time.After(timeout):
// 		fmt.Println("1")
// 		ch <- true
// 	}
// }
//
// func printTwo(ch chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("2")
//
// 	check := <-ch
// 	if check {
// 		fmt.Println("2")
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func boilPasta(done chan bool) {
	fmt.Println("Boiling pasta...")
	time.Sleep(2 * time.Second) // Simulate delay
	fmt.Println("Pasta is done!")
	done <- true // Notify that task is complete
}

func main() {
	fmt.Println("Starting tasks...")

	// Channels to sync tasks
	done1 := make(chan bool)
	done2 := make(chan bool)

	// Run tasks concurrently
	go boilPasta(done1)
	go boilPasta(done2)

	// Wait for tasks to complete
	<-done1
	<-done2

	fmt.Println("All tasks done!")
}
