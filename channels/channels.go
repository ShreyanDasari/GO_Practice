package main

import(
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("Received:", num)
		}
	}()

	wg.Wait()
	fmt.Println("All numbers processed!")
}