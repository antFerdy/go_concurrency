package main

import "fmt"

func main() {

	slce := []int{10, 20, 30}
	fmt.Println(cap(slce))

	old := slce
	slce = append(slce, 40)

	fmt.Println(cap(slce))
	fmt.Println(slce)
	slceP := &slce
	fmt.Print(slceP)

	fmt.Println(old)

}
