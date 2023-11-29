// https://leetcode.com/problems/word-search/
// 79. Word Search
//
// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
// Example 1:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true
//
// Example 2:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// Output: true
//
// Example 3:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// Output: false

package medium

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if dfs(board, i, j, word, 0) == true {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, row int, col int, word string, wordIdx int) bool {
	if wordIdx >= len(word) {
		return false
	}
	if (row >= len(board)) || (col >= len(board[0])) || (row < 0) || (col < 0) {
		return false
	}
	if word[wordIdx] != board[row][col] {
		return false
	}
	if wordIdx == (len(word) - 1) {
		return true
	}

	board[row][col] ^= byte(128)

	if (dfs(board, row+1, col, word, wordIdx+1) == true) ||
		(dfs(board, row-1, col, word, wordIdx+1) == true) ||
		(dfs(board, row, col+1, word, wordIdx+1) == true) ||
		(dfs(board, row, col-1, word, wordIdx+1) == true) {
		return true
	}

	board[row][col] ^= byte(128)

	return false
}
