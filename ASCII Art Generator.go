package main

import (
	"fmt"
	"strings"
)

var letters = map[rune][]string{
	'L': {
		"L     ",
		"L     ",
		"L     ",
		"L     ",
		"LLLLLL",
	},

	'I': {
		"IIIII",
		"  I  ",
		"  I  ",
		"  I  ",
		"IIIII",
	},

	'S': {
		" SSSS",
		"S    ",
		" SSS ",
		"    S",
		"SSSS ",
	},

	'N': {
		"N   N",
		"NN  N",
		"N N N",
		"N  NN",
		"N   N",
	},
}

func main() {
	var word string

	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	word = strings.ToUpper(word)

	// Print row by row
	for row := 0; row < 5; row++ {
		for _, ch := range word {
			if pattern, exists := letters[ch]; exists {
				fmt.Print(pattern[row], "  ")
			} else {
				fmt.Print("?????  ")
			}
		}
		fmt.Println()
	}
}
