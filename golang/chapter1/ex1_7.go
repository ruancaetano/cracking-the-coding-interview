package chapter1

import (
	"fmt"
)

// RunChapter1Exercise7
// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
// [
//
//	[1,2,3]
//	[4,5,6]
//	[7,8,9]
//
// ]
//
// [
//
//	[7,4,1]
//	[8,5,2]
//	[9,6,3]
//
// ]
func RunChapter1Exercise7() {
	fmt.Println("Original")
	fmt.Println([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println()

	fmt.Println("rotated 90")
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(rotateMatrix(input))
	rotateMatrixInPlace(input)
	fmt.Println(input)
	fmt.Println()

	fmt.Println("rotated 180")
	input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(rotateMatrix(rotateMatrix(input)))
	rotateMatrixInPlace(rotateMatrixInPlace(input))
	fmt.Println(input)
	fmt.Println()

	fmt.Println("rotated 270")
	input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(rotateMatrix(rotateMatrix(rotateMatrix(input))))
	rotateMatrixInPlace(rotateMatrixInPlace(rotateMatrixInPlace(input)))
	fmt.Println(input)
	fmt.Println()

	fmt.Println("rotated 360")
	input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(rotateMatrix(rotateMatrix(rotateMatrix(rotateMatrix(input)))))
	rotateMatrixInPlace(rotateMatrixInPlace(rotateMatrixInPlace(rotateMatrixInPlace(input))))
	fmt.Println(input)
	fmt.Println()

}

// rotate matrix with extra space
// Time:  O(N^2)
// Space:  O(N^2)
// Where N is the matrix column/row size
func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if result[j] == nil {
				result[j] = make([]int, n)
			}
			result[j][n-i-1] = matrix[i][j]
		}
	}
	return result
}

// rotate matrix without extra space
// Time: O(Nˆ2 + N^2/2) -> O(N^2)
// Space: O(1)
// Where N is the inputs length
func rotateMatrixInPlace(matrix [][]int) [][]int {
	n := len(matrix)

	// Tranversal
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			previous := matrix[j][i]    // 4
			matrix[j][i] = matrix[i][j] // 2
			matrix[i][j] = previous
		}
	}

	// Invert
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			previous := matrix[i][j]        // 1
			matrix[i][j] = matrix[i][n-j-1] // 7
			matrix[i][n-j-1] = previous
		}
	}

	return matrix
}
