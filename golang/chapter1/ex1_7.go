package chapter1

import (
	"fmt"
)

// Run1_7 Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
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
func Run1_7() {
	fmt.Println("Original")
	fmt.Println([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println("rotated 90")
	fmt.Println(rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println("rotated 180")
	fmt.Println(rotateMatrix(rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})))
	fmt.Println("rotated 270")
	fmt.Println(rotateMatrix(rotateMatrix(rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))))
	fmt.Println("rotated 360")
	fmt.Println(rotateMatrix(rotateMatrix(rotateMatrix(rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})))))

}

// rotateMatrix
// time O(N^2) where N is the matrix column/row size
// storage O(N) where N is the matrix size
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
