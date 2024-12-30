package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func B(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Routine B started")
	defer fmt.Println("Routine B stopped")

	wg.Add(1)
	go C(ctx, wg)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting B")
			return
		default:
			// fmt.Println("B is running...")
			time.Sleep(1 * time.Second)
		}
	}
}

