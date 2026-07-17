package main

import (
	"fmt"
	"unicode"
)

func checkPasswordStrength(password string) string {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	score := 0

	if len(password) >= 8 {
		score++
	}
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSpecial {
		score++
	}

	if score <= 2 {
		return "Weak"
	} else if score <= 4 {
		return "Medium"
	}

	return "Strong"
}

func main() {
	var password string

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	fmt.Println("Password:", password)
	fmt.Println("Password Strength:", checkPasswordStrength(password))
}
