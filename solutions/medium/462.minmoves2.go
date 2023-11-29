// 462. Minimum Moves to Equal Array Elements II
//
// Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
// In one move, you can increment or decrement an element of the array by 1.
// Test cases are designed so that the answer will fit in a 32-bit integer.
//
//Example 1:
// Input: nums = [1,2,3]
// Output: 2
// Explanation:
// Only two moves are needed (remember each move increments or decrements one element):
// [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
//
// Example 2:
// Input: nums = [1,10,2,9]
// Output: 16

package medium

import (
	"sort"
)

func minMoves2(nums []int) int {
	var avrg int
	var moves int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	halfLen := len(nums) / 2
	if len(nums)%2 == 0 {
		avrg = (nums[halfLen-1] + nums[halfLen]) / 2
	} else {
		avrg = nums[halfLen]
	}

	for _, n := range nums {
		if avrg > n {
			moves += avrg - n
		} else {
			moves += n - avrg
		}
	}

	return moves
}
