package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var i int
	go func() {
		for {
			i++
		}
	}()

	log.Println("i: ", i)
}
