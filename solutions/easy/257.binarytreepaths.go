// https://leetcode.com/problems/binary-tree-paths/
// 257. Binary Tree Paths
//
// Given the root of a binary tree, return all root-to-leaf paths in any order.
// A leaf is a node with no children.
//
// Example 1:
// Input: root = [1,2,3,null,5]
// Output: ["1->2->5","1->3"]
//
// Example 2:
// Input: root = [1]
// Output: ["1"]

package easy

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	printPath(root, strconv.Itoa(root.Val), &res)

	return res
}

func printPath(root *TreeNode, path string, res *[]string) {
	if (root.Left == nil) && (root.Right == nil) {
		*res = append(*res, path)

		return
	}

	if root.Left != nil {
		leftPath := path + "->" + strconv.Itoa(root.Left.Val)
		printPath(root.Left, leftPath, res)
	}
	if root.Right != nil {
		rightPath := path + "->" + strconv.Itoa(root.Right.Val)
		printPath(root.Right, rightPath, res)
	}
}
