package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	// Palindrome word
	word := "Madam"

	// Palindrome number
	number := 12321

	fmt.Println("Checking Word:", word)
	if isPalindrome(word) {
		fmt.Println(word, "is a palindrome.")
	} else {
		fmt.Println(word, "is not a palindrome.")
	}

	fmt.Println()

	fmt.Println("Checking Number:", number)
	if isPalindrome(strconv.Itoa(number)) {
		fmt.Println(number, "is a palindrome.")
	} else {
		fmt.Println(number, "is not a palindrome.")
	}
}
