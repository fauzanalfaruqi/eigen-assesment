package soal_2

func Longest(sentence string) string {
	var longestWord string
	var maxLen int
	var wordStart int

	for i, char := range sentence {
		if char == ' ' || i == len(sentence)-1 {
			wordLen := i - wordStart + 1
			if i == len(sentence)-1 {
				wordLen++
			}
			if wordLen > maxLen {
				longestWord = sentence[wordStart : i+1]
				maxLen = wordLen
			}
			wordStart = i + 1
		}
	}

	return longestWord
}
