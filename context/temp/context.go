package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

func contextTest() {
	printMemStats()
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, 1)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			if input == "q\n" {
				cancel()
				break
			}
		}
	}()
	fmt.Println("Press 'q' to stop the worker and exit")
	<-ctx.Done()
	printMemStats()
	fmt.Println("Exiting program")
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is shutting down\n", id)
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc: %v MB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc: %v MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys: %v MB\n", bToMb(m.Sys))
	fmt.Printf("NumGC: %v MB\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1000 / 1000
}
