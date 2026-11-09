package main

import (
	"fmt"
	"time"
)

func main() {
	goroutine_cycle()
}

func goroutine_cycle() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("\n new value of i %d", i)
		}()
	}

	time.Sleep(2 * time.Second)
}
