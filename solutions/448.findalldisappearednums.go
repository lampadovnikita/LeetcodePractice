// 448. Find All Numbers Disappeared in an Array
//
// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
//
// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
//
// Example 2:
// Input: nums = [1,1]
// Output: [2]

package main

func findDisappearedNumbers(nums []int) []int {
	var desiredIdx int
	for _, n := range nums {
		if n < 0 {
			desiredIdx = -n - 1
		} else {
			desiredIdx = n - 1
		}

		if nums[desiredIdx] > 0 {
			nums[desiredIdx] *= -1
		}
	}

	res := make([]int, 0)
	for i, n := range nums {
		if n > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
