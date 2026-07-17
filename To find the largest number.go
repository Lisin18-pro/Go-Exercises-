package main

import "fmt"

func main() {
	var n int

	// Ask the user how many numbers they want to enter
	fmt.Print("How many numbers do you want to enter? ")
	fmt.Scan(&n)

	// Check if the input is valid
	if n <= 0 {
		fmt.Println("Please enter a valid number of elements.")
		return
	}

	// Create a slice to store the numbers
	numbers := make([]int, n)

	// Take input from the user
	fmt.Println("Enter the numbers:")
	for i := 0; i < n; i++ {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	// Assume the first number is the largest
	largest := numbers[0]

	// Find the largest number
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}

	// Display the result
	fmt.Println("\nNumbers entered:", numbers)
	fmt.Println("Largest number:", largest)
}
