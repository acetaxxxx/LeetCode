package lc542

import (
	"math"
)

type calculateObj struct {
	I int
	J int
	M int
	N int
}

var verified map[int]bool
var readyToDo map[int]calculateObj
var onCalcu []bool

func updateMatrix(matrix [][]int) [][]int {

	verified = make(map[int]bool)
	readyToDo = make(map[int]calculateObj)
	m := len(matrix)
	n := len(matrix[0])
	onCalcu = make([]bool, m*n)
	t := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				verified[i*n+j] = true
			} else if verified[i*n+j] == false {
				readyToDo[t] = calculateObj{
					I: i,
					J: j,
					M: m,
					N: n,
				}
				t++
			}
		}
	}
	for t, value := range readyToDo {

		onCalcu = make([]bool, m*n)
		minDistance(value.I, value.J, value.M, value.N, matrix)
		delete(readyToDo, t)

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

	if allowed(i, j, m, n) && verified[i*n+j] == true {
		return matrix[i][j]
	}
	if onCalcu[i*n+j] == true {
		return math.MaxInt32 - 1
	}
	if matrix[i][j] == 0 {
		verified[i*n+j] = true
		return 0
	}
	minResult := math.MaxInt32
	onCalcu[i*n+j] = true
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
		if verified[i*n+j-1] == false {
			minResult = min(minResult, 1+minDistance(i, j-1, m, n, matrix))
		} else {
			minResult = min(minResult, 1+matrix[i][j-1])
		}

	}
	if allowed(i, j+1, m, n) {
		if verified[i*n+j+1] == false {
			minResult = min(minResult, 1+minDistance(i, j+1, m, n, matrix))
		} else {
			minResult = min(minResult, 1+matrix[i][j+1])
		}
	}
	verified[i*n+j] = true
	matrix[i][j] = minResult
	return minResult
}

func calresultmain(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j :=range result[i]{
			result[i][j] = math.MaxInt32
		}
	}
	
	

}
