// https://leetcode.com/problems/reverse-linked-list/
// 206. Reverse Linked List
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
// Example 2:
// Input: head = [1,2]
// Output: [2,1]
//
// Example 3:
// Input: head = []
// Output: []

package easy

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil

	for head != nil {
		next := head.Next

		head.Next = newHead

		newHead = head

		head = next
	}

	return newHead
}

func reverseListRecursive(head *ListNode) *ListNode {
	return reverseListRecursiveInner(head, nil)
}

func reverseListRecursiveInner(head *ListNode, newHead *ListNode) *ListNode {
	if head == nil {
		return newHead
	}

	next := head.Next
	head.Next = newHead

	return reverseListRecursiveInner(next, head)
}
