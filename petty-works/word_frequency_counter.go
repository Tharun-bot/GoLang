//PH1 - Write a hardcoded paragraph and store it in a var. Later user input
//split the words, count the number of word frequency with a map, print the top 5 words

package main

import (
	"fmt"
	"strings"
)

func main() {
	paragraph := `Write a program declaring every data type with their zero values. From memory after watching, not copy-paste.
Build a unit converter CLI: km↔miles, celsius↔fahrenheit, kg↔lbs. Pure stdlib, no packages. Use constants for conversion ratios.`

	fmt.Println("Splitted Words : ", strings.Split(paragraph, " "))

}
