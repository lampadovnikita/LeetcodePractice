// https://leetcode.com/problems/invert-binary-tree/
// 226. Invert Binary Tree
//
// Given the root of a binary tree, invert the tree, and return its root.
//
// Example 1:
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
//
// Example 2:
// Input: root = [2,1,3]
// Output: [2,3,1]
//
// Example 3:
// Input: root = []
// Output: []

package easy

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		currNode := queue[0]
		queue = queue[1:]

		tmp := currNode.Left
		currNode.Left = currNode.Right
		currNode.Right = tmp

		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}
	}

	return root
}
