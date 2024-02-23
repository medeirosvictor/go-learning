package main

import "fmt"

func main() {
	fmt.Println("Enter first number: ")

	var firstNumber int
	fmt.Scanln(&firstNumber)

	fmt.Println("Enter second number: ")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	fmt.Println("Enter operation: ")
	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println(firstNumber + secondNumber)
	case "-":
		fmt.Println(firstNumber - secondNumber)
	case "*":
		fmt.Println(firstNumber * secondNumber)
	case "/":
		if secondNumber == 0 {
			fmt.Println("Division by zero is not allowed")
			return
		}
		fmt.Println(firstNumber / secondNumber)
	default:
	}
}
