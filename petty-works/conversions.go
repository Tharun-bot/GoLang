//For mul, div. The operands must be of same type, Float32(x)
//fmt.Scanf(ref) - for user input

package main

import (
	"fmt"
	"math"
)

func main() {
	var celc int
	var km float32
	var kg float64

	fmt.Println("Enter celcius")
	fmt.Scan(&celc)
	fmt.Println("Enter KM")
	fmt.Scan(&km)
	fmt.Println("Enter KG")
	fmt.Scan(&kg)
	fahr := celcius_to_fahr(celc)
	fmt.Println("Fahr : ", fahr)
	miles := km_to_mi(km)
	fmt.Println("MIles : ", miles)
	lbs := kg_to_lbs(kg)
	fmt.Println("LBS : ", lbs)
}

func celcius_to_fahr(celc int) float64 {
	c := float64(celc)
	const conv = 1.8
	return math.Round((c * conv) + 32)
}

func km_to_mi(km float32) float32 {
	return km * 0.62
}

func kg_to_lbs(kg float64) float64 {
	return kg * 2.204
}
