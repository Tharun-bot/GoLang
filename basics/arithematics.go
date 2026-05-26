package main

import (
	"fmt"
	"math"
)

func main() {
	//variable decleration
	var a, b int = 12, 10
	var result int

	result = a + b
	fmt.Println("Addition : ", result)

	result = a - b
	fmt.Println("Subtraction : ", result)

	result = a * b
	fmt.Println("MUltiplication : ", result)

	result = a / b
	fmt.Println("Division : ", result)

	//overflow signed with integers

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
