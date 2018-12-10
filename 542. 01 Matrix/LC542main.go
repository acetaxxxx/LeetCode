package lc542

var verified map[int]bool

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	print(verified[1] == false)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				verified[i*m+j] = true
			} else if verified[i*m+n] == false {
				matrix[i][j] = minDistance(1, 2, 3, 4, matrix)
			}
		}
	}
	// verified := make([][]bool,2,false)
	return [][]int{{1, 2, 3}, {4, 5, 6}}
}

func allowed(i, j, m, n int) bool {
	return i > 0 && j > 0 && i < m && j < n
}

func minDistance(i, j, m, n int, matrix [][]int) int {

	return 1
}
