package soal_4

import "math"

func DiagonalDifference(matrix [][]int) int {
	n := len(matrix)
	var diagonal1, diagonal2 int

	for i := 0; i < n; i++ {
		diagonal1 += matrix[i][i]
		diagonal2 += matrix[i][n-i-1]
	}

	return int(math.Abs(float64(diagonal1 - diagonal2)))
}
