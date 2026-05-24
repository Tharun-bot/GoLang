package main

import (
	"fmt"
)

func main() {
	add := adder()
	fmt.Println("Sum : ", add())
	fmt.Println("Sum : ", add())
	fmt.Println("Sum : ", add())
	fmt.Println("Sum : ", add())

	subtractor := func() func(x int) int {
		countdown := 100
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println("Coundown : ", subtractor(1))
	fmt.Println("Coundown : ", subtractor(1))
	fmt.Println("Coundown : ", subtractor(1))
	fmt.Println("Coundown : ", subtractor(1))
}

func adder() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}
