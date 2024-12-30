package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func D(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Routine D started")
	defer fmt.Println("Routine D stopped")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting D")
			return
		default:
			// fmt.Println("D is running...")
			time.Sleep(1 * time.Second)
		}
	}
}
