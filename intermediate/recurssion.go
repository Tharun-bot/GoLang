package intermediate

import "fmt"

func main() {
	fmt.Println("Factorial : ", factorial(6))
	fmt.Println("Sum of digits : ", sumOfDigits(987))
}

func factorial(x int) int {
	if x == 2 {
		return 2
	}
	fact := x
	fact = fact * factorial(x-1)
	return fact
}

func sumOfDigits(x int) int {
	if x < 10 {
		return x
	}

	return x%10 + sumOfDigits(x/10)
}
