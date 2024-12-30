package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go A(ctx, &wg)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press 'q' to stop all routines and exit.")
		for {
			input, _ := reader.ReadString('\n')
			if input == "q\n" {
				fmt.Println("Canceling context and stopping routines...")
				cancel()
				return
			}
		}
	}()

	fmt.Printf("Before wait(), ACTIVE routines are: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("After wait(), ACTIVE routines are: %d\n", runtime.NumGoroutine())
	fmt.Println("All routines stopped. Exiting Program")
}
