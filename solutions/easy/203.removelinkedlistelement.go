// https://leetcode.com/problems/remove-linked-list-elements/
// 203. Remove Linked List Elements
//
// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.
//
// Example 1:
// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]
//
// Example 2:
//
// Input: head = [], val = 1
// Output: []
//
// Example 3:
// Input: head = [7,7,7,7], val = 7
// Output: []

package easy

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	curr := dummy

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}
