package main

import "fmt"

func main() {
	var word string
	var reversed string

	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	// Reverse the word
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	fmt.Println("Reversed word:", reversed)
}
