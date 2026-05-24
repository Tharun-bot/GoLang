package basics

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	guess_answer := rand.IntN(100)
	fmt.Println("Guess answer is :", guess_answer)

	is_correct := false
	attempt := 0

	for {
		if is_correct {
			fmt.Printf("Your guess is correct with %d attempts\n", attempt)
			break
		}

		var input int
		fmt.Scan(&input)

		if input == guess_answer {
			is_correct = true
			continue
		} else if input > guess_answer {
			fmt.Println("Too big")
			attempt++
		} else {
			fmt.Println("Too Small")
			attempt++
		}
	}

}
