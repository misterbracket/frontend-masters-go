package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int
	var result float64

	fmt.Println("Enter the operation you want to perform: ")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	// read the input from user
	fmt.Scanln(&operation)

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&number1)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&number2)

	switch operation {
	case "1":
		result = float64(number1 + number2)
	case "2":
		result = float64(number1 - number2)
	case "3":
		result = float64(number1 * number2)
	case "4":
		result = float64(number1) / float64(number2)
	}

	fmt.Println("Result is: ", result)
}
