package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		return 0
	}
	return a / b
}

func main() {
	var num1, num2 float64
	var operator string
	var result float64

	// Take input from the user
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	// Perform the calculation
	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		fmt.Println("Invalid operator")
		return
	}

	// Display the result
	fmt.Println("\n------ Result ------")
	fmt.Println("First Number :", num1)
	fmt.Println("Operator     :", operator)
	fmt.Println("Second Number:", num2)
	fmt.Println("Result       :", result)
}
