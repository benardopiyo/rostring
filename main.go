package main

import (
	"os"
	"github.com/01-edu/z01"
)

// splitWords splits a string into words based on spaces.
func splitWords(s string) []string {
	words := []string{}
	start := 0
	inWord := false

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' {
			if !inWord {
				start = i
				inWord = true
			}
		} else {
			if inWord {
				words = append(words, s[start:i])
				inWord = false
			}
		}
	}
	if inWord {
		words = append(words, s[start:])
	}

	return words
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	words := splitWords(os.Args[1])

	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Rotate words
	firstWord := words[0]
	rotatedWords := append(words[1:], firstWord)

	// Print rotated words
	for i, word := range rotatedWords {
		if i > 0 {
			z01.PrintRune(' ')
		}
		for _, r := range word {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
