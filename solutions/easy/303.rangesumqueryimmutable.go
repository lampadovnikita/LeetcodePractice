// https://leetcode.com/problems/range-sum-query-immutable/
// 303. Range Sum Query - Immutable
//
// Given an integer array nums, handle multiple queries of the following type:
// Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the NumArray class:
// NumArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
//
// Example 1:
// Input
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// Output
// [null, 1, -1, -3]
// Explanation
// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
// numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
// numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

package easy

type NumArray struct {
	nums []int

	// index -> sum
	table []int
}

func Constructor(nums []int) NumArray {
	table := make([]int, len(nums))

	sum := 0
	for idx := range nums {
		sum += nums[idx]

		table[idx] = sum
	}

	return NumArray{
		nums:  nums,
		table: table,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.table[right]
	}

	return this.table[right] - this.table[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
