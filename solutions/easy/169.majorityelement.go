// https://leetcode.com/problems/majority-element/
// 169. Majority Element
//
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

package easy

// Boyer–Moore majority vote algorithm
func majorityElement(nums []int) int {
	var maj int

	counter := 0
	for i := 0; i < len(nums); i++ {
		if counter == 0 {
			maj = nums[i]
			counter = 1
		} else if maj == nums[i] {
			counter++
		} else {
			counter--
		}
	}

	return maj
}
