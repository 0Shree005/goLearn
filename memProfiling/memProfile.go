package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Memory stats before:")
	printMemStats()

	s := make([]int, 10_000_000)
	for i := range s {
		s[i] = i
	}
	fmt.Println()

	fmt.Println("Memory stats after:")
	printMemStats()

	fmt.Println()

	runtime.GC()

	fmt.Println("Memory stats after GC:")
	printMemStats()

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
