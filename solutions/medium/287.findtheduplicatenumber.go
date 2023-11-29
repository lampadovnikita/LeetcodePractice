// https://leetcode.com/problems/find-the-duplicate-number/
// 287. Find the Duplicate Number
//
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.
//
// Example 1:
// Input: nums = [1,3,4,2,2]
// Output: 2
//
// Example 2:
// Input: nums = [3,1,3,4,2]
// Output: 3

package medium

// This problem is as same as 142. Linked List Cycle II,
// you can refer to this solution for the explanation of the slow fast pointer approach to solve this problem.
func findDuplicate(nums []int) int {
	slow := nums[nums[0]]
	fast := nums[nums[nums[0]]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
