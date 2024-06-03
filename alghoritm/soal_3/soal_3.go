package soal_3

func CountOccurrences(INPUT, QUERY []string) []int {
	wordCount := make(map[string]int)
	for _, word := range INPUT {
		wordCount[word]++
	}

	var result []int
	for _, word := range QUERY {
		result = append(result, wordCount[word])
	}
	return result
}
