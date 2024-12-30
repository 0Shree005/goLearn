package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func C(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Routine C started")
	defer fmt.Println("Routine C stopped")

	wg.Add(1)
	go D(ctx, wg)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting C")
			return
		default:
			// fmt.Println("C is running...")
			time.Sleep(1 * time.Second)
		}
	}
}
