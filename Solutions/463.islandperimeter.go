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

package main

func islandPerimeter(grid [][]int) int {
	var per int

	for row, lay := range grid {
		for col, surf := range lay {
			if surf == 1 {
				per += getBord(grid, col, row, 0, -1)
				per += getBord(grid, col, row, 1, 0)
				per += getBord(grid, col, row, 0, 1)
				per += getBord(grid, col, row, -1, 0)
			}
		}
	}

	return per
}

func getBord(grid [][]int, xCur int, yCur int, xBias int, yBias int) int {
	xTar := xCur + xBias
	if xTar < 0 {
		return 1
	} else if xTar >= len(grid[0]) {
		return 1
	}

	yTar := yCur + yBias
	if yTar < 0 {
		return 1
	} else if yTar >= len(grid) {
		return 1
	}

	if grid[yTar][xTar] == 1 {
		return 0
	} else {
		return 1
	}
}
