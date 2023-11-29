// https://leetcode.com/problems/island-perimeter/
// 463. Island Perimeter
//
// You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.
// Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
// and there is exactly one island (i.e., one or more connected land cells).
// The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1.
// The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
//
// Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
// Output: 16
// Explanation: The perimeter is the 16 yellow stripes in the image above.

package easy

func islandPerimeter(grid [][]int) int {
	lands := 0
	neighbourings := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				lands += 1

				if (row < len(grid)-1) && (grid[row+1][col] == 1) {
					neighbourings += 1
				}
				if (col < len(grid[row])-1) && (grid[row][col+1] == 1) {
					neighbourings += 1
				}
			}
		}
	}

	return 4*lands - 2*neighbourings
}
