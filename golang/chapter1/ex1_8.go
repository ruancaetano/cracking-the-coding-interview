package chapter1

import (
	"fmt"
)

// RunChapter1Exercise8
// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.
//
// Input
// [
//
//	[1,2,3]
//	[4,5,6]
//	[7,8,0]
//
// ]
//
// Output
// [
//
//	[1,2,0]
//	[4,5,0]
//	[0,0,0]
//
// ]
func RunChapter1Exercise8() {
	fmt.Println(solveZeroMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}))
}

// solveZeroMatrix
// Time: O(NxM)
// Space: O(N+M)
func solveZeroMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])

	zerosRows := make([]bool, n)
	zeroColumns := make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zerosRows[i] = true
				zeroColumns[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if zerosRows[i] || zeroColumns[j] {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}
