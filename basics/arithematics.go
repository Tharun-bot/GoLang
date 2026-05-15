package main

import "fmt"

func main() {
	var a, b int = 1, 2
	var sum = a + b

	c, d := 1.0, 2.0

	var div float64 = c / d

	fmt.Printf("Sum is : %d\n", sum)
	fmt.Println("Division : ", div)

}
