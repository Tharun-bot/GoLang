package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message1 := "Hello \nWorld!"
	message2 := "Hello \tWorld"
	rawMessage := `Hello\nWorld`

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message1 : ", len(message1))
	fmt.Println("Length of rawMessage : ", len(rawMessage))

	fmt.Println("First char of message1 is : ", message1[0]) // gives the ascii character
	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for index, char := range message1 {
		fmt.Printf("Index %d and char %v\n", index, char) //%v - gives ASCII Values and %c gives the actual char and %d gives the integer
	}

	fmt.Println("Rune count : ", utf8.RuneCountInString(message1))

}
