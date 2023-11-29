// https://leetcode.com/problems/average-of-levels-in-binary-tree/
// 637. Average of Levels in Binary Tree
//
// Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: [3.00000,14.50000,11.00000]
// Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
// Hence return [3, 14.5, 11].
//
// Example 2:
// Input: root = [3,9,20,15,7]
// Output: [3.00000,14.50000,11.00000]

package easy

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var currLevelDivider float64 = 1.0
	var currLevelAvrg float64 = 0.0

	currLevelLen := 1
	nextLevelLen := 0
	currLevelSum := 0

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if currNode.Left != nil {
			nextLevelLen++
			queue = append(queue, currNode.Left)
		}
		if currNode.Right != nil {
			nextLevelLen++
			queue = append(queue, currNode.Right)
		}

		currLevelSum += currNode.Val
		currLevelLen--

		if currLevelLen == 0 {
			currLevelAvrg = float64(currLevelSum) / currLevelDivider
			res = append(res, currLevelAvrg)

			currLevelLen = nextLevelLen
			currLevelDivider = float64(nextLevelLen)
			nextLevelLen = 0
			currLevelSum = 0
		}
	}

	return res
}
