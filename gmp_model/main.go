package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of P", runtime.GOMAXPROCS(0))
	fmt.Println("Number of CPU", runtime.NumCPU())

	runtime.GOMAXPROCS(16)

	fmt.Println("AFTER")
	fmt.Println("Number of P", runtime.GOMAXPROCS(0))
	fmt.Println("Number of CPU", runtime.NumCPU())
}
