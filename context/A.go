package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func A(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Routine A started")
	defer fmt.Println("Routine A stopped")

	wg.Add(1)
	go B(ctx, wg)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting A")
			return
		default:
			// fmt.Println("A is running...")
			time.Sleep(1 * time.Second)
		}
	}
}
