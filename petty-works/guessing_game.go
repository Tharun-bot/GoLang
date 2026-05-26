package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fixed_random_number := rand.IntN(100)

	var is_correct bool = false
	count := 0

	for !is_correct {
		var guess int32
		fmt.Println("Enter your guess : ")
		fmt.Scanln(&guess)
		if guess > 100 || guess < 0 {
			fmt.Println("Input out of range")
			continue
		}

		if guess == int32(fixed_random_number) {
			count += 1
			fmt.Printf("Success in %d attempts\n", count)
			is_correct = true
			break
		} else if guess > int32(fixed_random_number) {
			fmt.Println("Number too big")
			count += 1
			continue
		} else {
			fmt.Println("Number too small")
			count += 1
			continue
		}

	}
}
