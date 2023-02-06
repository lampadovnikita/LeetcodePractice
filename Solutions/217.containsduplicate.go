// 217. Contains Duplicate
//
// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
//
// Example 1:
// Input: nums = [1,2,3,1]
// Output: true
//
// Example 2:
// Input: nums = [1,2,3,4]
// Output: false
//
// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

package main

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    
    for _, n := range nums {
        _, ok := m[n]
        if ok == false {
            m[n] = false
        } else {
            return true
        }
    }
    
    return false
}