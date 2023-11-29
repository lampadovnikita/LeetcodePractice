// https://leetcode.com/problems/palindrome-linked-list/
// 234. Palindrome Linked List
//
// Given the head of a singly linked list, return true if it is a
// palindrome or false otherwise.
//
// Example 1:
// Input: head = [1,2,2,1]
// Output: true
//
// Example 2:
// Input: head = [1,2]
// Output: false

package easy

func isPalindrome(head *ListNode) bool {
	middle := head
	fast := head

	for fast != nil && fast.Next != nil {
		middle = middle.Next

		fast = fast.Next.Next
	}

	newHead := middle
	curr := middle.Next

	for curr != nil {
		next := curr.Next

		curr.Next = newHead

		newHead = curr
		curr = next
	}

	left := head
	right := newHead

	for left != middle {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}
