package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

var totalFilesUpdated = 0;

func processFile(id int) {
    fmt.Printf("Processing file %d...\n", id)
    time.Sleep(time.Second) // Simulate work
	mu.Lock()
	totalFilesUpdated++
	mu.Unlock()
    fmt.Printf("File %d done!\n", id)
}

func main() {
	wg.Add(1)
    for i := 1; i <= 1000; i++ {
       go func(id int) {
           processFile(id)
		   defer wg.Done()
       }(i)
    }
	fmt.Println("All files are being processed...")
	wg.Wait()
	fmt.Println("Total files updated:", totalFilesUpdated)
	fmt.Println("All files processed!")
   
}