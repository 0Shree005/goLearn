package fileChange

import (
	"fmt"
	"time"
)

func FileChange(mainChan chan string) {
	fmt.Println("fileChange")
	ticker := time.NewTicker(1 * time.Second)
	mainSendTicker := time.NewTicker(3 * time.Second)

	fileNameChan := make(chan string)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-mainSendTicker.C:
				fileNameRes := sendToClient(fileNameChan)
				mainChan <- fileNameRes
			}
		}
	}()
}

func sendToClient(fileNameChan chan string) string {
	fileName := ""
	for fileName := range fileNameChan {
		return fileName
	}
	return fileName
}
