package main

import (
	"log"
	"time"
)

func main() {
	launchNeverExit(task)
	launchNeverExit(task)
}

func task() {
	log.Println("Executing task")
	time.Sleep(3 * time.Second)
	panic("Aaahhhhhhh!")
}

func launchNeverExit(task func()) {
	defer func() {
		v := recover()
		if v != nil {
			log.Printf("Recovering from panic %v", v)
			go launchNeverExit(task)
		}
	}()

	if task != nil {
		task()
	}
}
