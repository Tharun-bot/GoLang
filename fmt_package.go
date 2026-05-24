package main

import "fmt"

func main() {
	// print("Hello")
	// print(12, 333)

	fmt.Println(12.00)

	name := "Tharun"
	age := 21

	fmt.Printf("Name : %s and age is %d\n", name, age)
	fmt.Printf("Binary : %b and hex is %X\n", name, name)

	//Formatting functions
	s := fmt.Sprint("Hello", "nonsense", 12000)
	fmt.Println("Sprint : ", s)

	sf := fmt.Sprintf("Hello", "Nonsense")
	fmt.Print("SF : ", sf)

}
