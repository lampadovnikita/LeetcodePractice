// https://leetcode.com/problems/path-sum/
// 112. Path Sum
//
// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
// A leaf is a node with no children.
//
// Example 1:
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation: The root-to-leaf path with the target sum is shown.

package easy

type SumNode struct {
	node       *TreeNode
	parentsSum int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := make([]SumNode, 0)
	stack = append(stack, SumNode{
		node:       root,
		parentsSum: 0,
	})

	for len(stack) != 0 {
		sNode := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if sNode.node.Left != nil {
			stack = append(stack, SumNode{
				node:       sNode.node.Left,
				parentsSum: sNode.parentsSum + sNode.node.Val,
			})
		}
		if sNode.node.Right != nil {
			stack = append(stack, SumNode{
				node:       sNode.node.Right,
				parentsSum: sNode.parentsSum + sNode.node.Val,
			})
		}

		if (sNode.node.Left == nil) && (sNode.node.Right == nil) {
			if (sNode.parentsSum + sNode.node.Val) == targetSum {
				return true
			}
		}
	}

	return false
}

func hasPathSumRecursive(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if (root.Left == nil) && (root.Right == nil) && ((targetSum - root.Val) == 0) {
		return true
	}

	return hasPathSumRecursive(root.Left, targetSum-root.Val) || hasPathSumRecursive(root.Right, targetSum-root.Val)
}
