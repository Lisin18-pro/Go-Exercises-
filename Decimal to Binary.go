package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&num)

	fmt.Printf("Binary: %b\n", num)
}
