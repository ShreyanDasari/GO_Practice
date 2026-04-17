package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func process(i int){
	fmt.Println("Processing", i)
	time.Sleep(5 * time.Second)
	fmt.Println("Done processing", i)
}
 func main() {
	for i:= 1; i<=100; i++ {
		wg.Add(1)
		go func(i int) {
			process(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("All processes done")
 }