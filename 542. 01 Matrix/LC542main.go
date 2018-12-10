package lc542

import (
	"math"
)

var verified map[int]bool

func updateMatrix(matrix [][]int) [][]int {
	verified = make(map[int]bool)
	m := len(matrix)
	n := len(matrix[0])
	print(verified[1] == false)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				verified[i*m+j] = true
			} else if verified[i*m+j] == false {
				matrix[i][j] = minDistance(i, j, m, n, matrix)
			}
		}
	}
	// verified := make([][]bool,2,false)
	return matrix
}

func allowed(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func minDistance(i, j, m, n int, matrix [][]int) int {

	if matrix[i][j] == 0 {
		verified[i*m+j] = true
		return 0
	}
	minResult := math.MaxInt32
	if allowed(i-1, j, m, n) {
		if verified[(i-1)*m+j] == false {
			minResult = min(minResult, 1+minDistance(i-1, j, m, n, matrix))
		} else {
			minResult = min(minResult, 1+matrix[i-1][j])
		}

	}
	if allowed(i+1, j, m, n) {
		if verified[(i+1)*m+j] == false {
			minResult = min(minResult, 1+minDistance(i+1, j, m, n, matrix))
		} else {
			minResult = min(minResult, 1+matrix[i+1][j])
		}

	}
	if allowed(i, j-1, m, n) {
		if verified[i*m+j-1] == false {
			minResult = min(minResult, 1+minDistance(i, j-1, m, n, matrix))
		} else {
			minResult = min(minResult, 1+matrix[i][j-1])
		}

	}
	if allowed(i, j+1, m, n) {
		if verified[i*m+j+1] == false {
			minResult = min(minResult, 1+minDistance(i, j+1, m, n, matrix))
		} else {
			minResult = min(minResult, 1+matrix[i][j+1])
		}
	}
	return minResult
}
