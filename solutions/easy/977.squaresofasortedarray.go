// https://leetcode.com/problems/squares-of-a-sorted-array/
// 977. Squares of a Sorted Array
//
// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
//
// Example 1:
// Input: nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
// Explanation: After squaring, the array becomes [16,1,0,9,100].
// After sorting, it becomes [0,1,9,16,100].
//
// Example 2:
// Input: nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]

package easy

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	left := 0
	right := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
	}

	return res
}

func abs(val int) int {
	if val > 0 {
		return val
	} else {
		return -val
	}
}
