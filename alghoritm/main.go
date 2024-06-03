package main

import (
	"alghoritm/soal_1"
	"alghoritm/soal_2"
	"alghoritm/soal_3"
	"alghoritm/soal_4"

	"fmt"
)

func main() {
	// Soal 1
	word := "NEGIE1"

	reversedWord := soal_1.ReverseWordWithNumber(word)
	fmt.Println(reversedWord)

	// Soal 2
	const sentence = "Saya sangat senang mengerjakan soal algoritma"

	longestWord := soal_2.Longest(sentence)
	fmt.Printf("%s\n", longestWord)

	// Soal 3
	INPUT := []string{"xc", "dz", "bbb", "dz"}
	QUERY := []string{"bbb", "ac", "dz"}

	output := soal_3.CountOccurrences(INPUT, QUERY)
	fmt.Println(output)

	// Soal 4
	matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}

	result := soal_4.DiagonalDifference(matrix)
	fmt.Println(result)
}
