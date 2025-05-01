package main

import "fmt"

func main() {
	// Print a simple message
	fmt.Println("Hello, World!")

	// Printing different types of values

	// Integers
	number := 42
	fmt.Println("Number:", number)

	// Floating point
	decimal := 3.14
	fmt.Println("Decimal:", decimal)

	// String
	name := "Gopher"
	fmt.Println("Name:", name)

	// Boolean
	isGolangFun := true
	fmt.Println("Is Golang fun?", isGolangFun)

	// Formatted printing
	fmt.Printf("The value of number is %d and name is %s\n", number, name)
}
