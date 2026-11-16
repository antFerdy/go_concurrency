package main

import (
	"log"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			counter++
		}()
	}

	wg.Wait()
	log.Println("Counter value", counter)

}
