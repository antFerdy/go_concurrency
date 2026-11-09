package main

import (
	"log"
	"time"
)

func main() {
	defer func() {
		v := recover()
		if v != nil {
			log.Println("recovering")
		}
	}()

	go func() {
		log.Println("Lauching new routine")
		panic("aaaaahhh!")
	}()

	time.Sleep(2 * time.Second)
}
