// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// 421. Maximum XOR of Two Numbers in an Array
//
// Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 <= i <= j < n.
//
// Example 1:
// Input: nums = [3,10,5,25,2,8]
// Output: 28
// Explanation: The maximum result is 5 XOR 25 = 28.
//
// Example 2:
// Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
// Output: 127

package medium

import "math"

type Trie struct {
	childs [2]*Trie
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(num int) {
	mask := int(math.Pow(2, 30))

	currTrie := t
	currIdx := 0

	for i := 0; i < 31; i++ {
		if (num & mask) == 0 {
			currIdx = 0
		} else {
			currIdx = 1
		}

		if currTrie.childs[currIdx] == nil {
			currTrie.childs[currIdx] = NewTrie()
		}

		currTrie = currTrie.childs[currIdx]

		mask = mask >> 1
	}
}

func (t *Trie) MaximumXOR(num int) int {
	currTrie := t

	mask := int(math.Pow(2, 30))

	res := 0
	currDigit := 0
	preferredDigit := 1

	for i := 0; i < 31; i++ {
		if (num & mask) == 0 {
			currDigit = 0
			preferredDigit = 1
		} else {
			currDigit = 1
			preferredDigit = 0
		}

		if currTrie.childs[preferredDigit] == nil {
			currTrie = currTrie.childs[currDigit]
		} else {
			res += int(math.Pow(2, 31-float64(i)-1))

			currTrie = currTrie.childs[preferredDigit]
		}

		mask = mask >> 1
	}

	return res
}

func findMaximumXOR(nums []int) int {
	t := NewTrie()

	for i := range nums {
		t.Insert(nums[i])
	}

	max := -1

	for i := range nums {
		xorRes := t.MaximumXOR(nums[i])
		if xorRes > max {
			max = xorRes
		}
	}

	return max
}
