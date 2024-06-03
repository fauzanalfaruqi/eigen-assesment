package soal_1

import (
	"unicode"
)

func ReverseWordWithNumber(word string) string {
	var letters, numbers string

	for _, char := range word {
		if unicode.IsLetter(char) {
			letters += string(char)
		} else if unicode.IsNumber(char) {
			numbers += string(char)
		}
	}

	reversedLetters := reverseString(letters)

	reversedWord := reversedLetters + numbers

	return reversedWord
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
