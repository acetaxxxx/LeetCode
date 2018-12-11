package lc542

import (
	"math"
	"github.com/golang-collections/collections/stack"
)

var verified map[int]bool
var readyToDo *stack.Stack

type calculateObj struct {
	I int
	J int
	M int 
	N int	
}

func updateMatrix(matrix [][]int) [][]int {
	verified = make(map[int]bool)
	readyToDo = stack.New()
	m := len(matrix)
	n := len(matrix[0])
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				verified[i*m+j] = true
			} else if verified[i*m+j] == false {	
				readyToDo.Push(calculateObj{
					I:i,
					J:j,
					M:m,
					N:n,					
				})						
			}
		}
	}
	for ;readyToDo.Len()>0;{
		
		tmp :=  readyToDo.Pop().(calculateObj)  
		minDistance(tmp.I,tmp.J,tmp.M,tmp.N,matrix)
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
	
	if allowed(i,j,m,n) && verified[i*m+j]==true {
		return matrix[i][j]
	}
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
	verified[i*m+j] = true
	matrix[i][j] = minResult
	return minResult
}
