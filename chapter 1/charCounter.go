package main

import (
	"fmt"
	"unicode"
)

func countCharacters(text string) map[string]int {
	result := map[string]int{
		"word_count":    0,
		"space_count":   0,
		"special_count": 0,
	}

	inWord := false
	for _, c := range text {
		if unicode.IsSpace(c) {
			result["space_count"]++
			inWord = false
		} else if unicode.IsLetter(c) || unicode.IsNumber(c) {
			if !inWord {
				result["word_count"]++
				inWord = true
			}
		} else {
			result["special_count"]++
		}
	}
	return result
}

func main() {
	text := "Merhaba, bu bir test mesajıdır! Kelime sayısı doğru mu?"
	counts := countCharacters(text)
	fmt.Println("Karakter Sayısı:", counts)
}
