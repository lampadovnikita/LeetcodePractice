// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// 83. Remove Duplicates from Sorted List
//
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
//
// Example 1:
// Input: head = [1,1,2]
// Output: [1,2]
//
// Example 2:
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]

package easy

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head

	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
