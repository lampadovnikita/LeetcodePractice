// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// 104. Maximum Depth of Binary Tree
//
// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
//
// Example 2:
// Input: root = [1,null,2]
// Output: 2

package easy

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	maxDepth := 0

	for len(queue) != 0 {
		currLevelLen := len(queue)

		for i := 0; i < currLevelLen; i++ {
			currNode := queue[0]
			queue = queue[1:]

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}

			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}

		maxDepth++
	}

	return maxDepth
}
