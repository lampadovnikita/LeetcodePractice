// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// 111. Minimum Depth of Binary Tree
//
// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
// Note: A leaf is a node with no children.
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 2
//
// Example 2:
// Input: root = [2,null,3,null,4,null,5,null,6]
// Output: 5

package easy

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	nextLevelLen := 0
	levelLen := 1
	minLevel := 1

	var currNode *TreeNode
	for len(queue) != 0 {
		currNode = queue[0]
		queue = queue[1:]

		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
			nextLevelLen++
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
			nextLevelLen++
		}

		if (currNode.Left == nil) && (currNode.Right == nil) {
			return minLevel
		}

		levelLen--

		if levelLen == 0 {
			minLevel++

			levelLen = nextLevelLen
			nextLevelLen = 0
		}
	}

	return minLevel
}
