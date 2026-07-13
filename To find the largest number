//Every Go program starts with a package declaration.
main is a special package that tells Go this is an executable program
package main 

//provides functions for printing output to the console
import "fmt"

func main() {
	// Sample numbers
	numbers := []int{10, 5, 8, 20, 3}

	// Assume the first number is the largest
	largest := numbers[0]

	// Find the largest number
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Largest number:", largest)
}
