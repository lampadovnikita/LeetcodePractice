// https://leetcode.com/problems/merge-two-binary-trees/
// 617. Merge Two Binary Trees
// You are given two binary trees root1 and root2.
// Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.
// Return the merged tree.
// Note: The merging process must start from the root nodes of both trees.
//
// Example 1:
// Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// Output: [3,4,5,5,4,null,7]
//
// Example 2:
// Input: root1 = [1], root2 = [1,2]
// Output: [2,2]

package easy

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if (root1 == nil) && (root2 == nil) {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue1 := make([]*TreeNode, 0)
	queue1 = append(queue1, root1)

	queue2 := make([]*TreeNode, 0)
	queue2 = append(queue2, root2)

	for len(queue1) != 0 {
		node1 := queue1[0]
		queue1 = queue1[1:]

		node2 := queue2[0]
		queue2 = queue2[1:]

		node1.Val += node2.Val

		if (node1.Left == nil) && (node2.Left != nil) {
			node1.Left = node2.Left
		} else if (node1.Left != nil) && (node2.Left != nil) {
			queue1 = append(queue1, node1.Left)
			queue2 = append(queue2, node2.Left)
		}

		if (node1.Right == nil) && (node2.Right != nil) {
			node1.Right = node2.Right
		} else if (node1.Right != nil) && (node2.Right != nil) {
			queue1 = append(queue1, node1.Right)
			queue2 = append(queue2, node2.Right)
		}
	}

	return root1
}

func mergeTreesRecursive(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if (root1 == nil) && (root2 == nil) {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Left = mergeTreesRecursive(root1.Left, root2.Left)
	root1.Right = mergeTreesRecursive(root1.Right, root2.Right)

	root1.Val += root2.Val

	return root1
}
