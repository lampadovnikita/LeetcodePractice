// https://leetcode.com/problems/spiral-matrix/
// 54. Spiral Matrix
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

package medium

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0, len(matrix)*len(matrix[0]))

	rowStart := 0
	rowEnd := len(matrix) - 1

	colStart := 0
	colEnd := len(matrix[0]) - 1

	for (rowStart <= rowEnd) && (colStart <= colEnd) {
		for j := colStart; j <= colEnd; j++ {
			res = append(res, matrix[rowStart][j])
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--

		if rowStart <= rowEnd {
			for j := colEnd; j >= colStart; j-- {
				res = append(res, matrix[rowEnd][j])
			}
		}
		rowEnd--

		if colStart <= colEnd {
			for i := rowEnd; i >= rowStart; i-- {
				res = append(res, matrix[i][colStart])
			}
		}
		colStart++
	}

	return res
}
