// https://leetcode.com/problems/find-all-duplicates-in-an-array/
// 442. Find All Duplicates in an Array
//
// Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.
// You must write an algorithm that runs in O(n) time and uses only constant extra space.
//
// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [2,3]
//
// Example 2:
// Input: nums = [1,1,2]
// Output: [1]
//
// Example 3:
// Input: nums = [1]
// Output: []

package medium

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := range nums {
		idx := abs(nums[i]) - 1

		if nums[idx] < 0 {
			res = append(res, abs(nums[i]))
		} else {
			nums[idx] *= -1
		}
	}

	return res
}

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
