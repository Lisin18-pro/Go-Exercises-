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
	return a / b
}

func main() {

	num1 := 0.0
	num2 := .0
	operator := "+"

	var result float64

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

	fmt.Println("First number:", num1)
	fmt.Println("Operator:", operator)
	fmt.Println("Second number:", num2)
	fmt.Println("Result:", result)
}
