// https://leetcode.com/problems/set-matrix-zeroes/
// 73. Set Matrix Zeroes
//
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
// You must do it in place.
//
// Example 1:
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
//
// Example 2:
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

package medium

func setZeroes(matrix [][]int) {
	zeroCols := false
	zeroRows := false

	for j := range matrix[0] {
		if matrix[0][j] == 0 {
			zeroCols = true
		}
	}
	for i := range matrix {
		if matrix[i][0] == 0 {
			zeroRows = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if (matrix[i][0] == 0) ||
				(matrix[0][j] == 0) {
				matrix[i][j] = 0
			}
		}
	}

	if zeroCols == true {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
	if zeroRows == true {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
