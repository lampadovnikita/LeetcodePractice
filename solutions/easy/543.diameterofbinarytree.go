// https://leetcode.com/problems/diameter-of-binary-tree/
// 543. Diameter of Binary Tree
//
// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
// The length of a path between two nodes is represented by the number of edges between them.
//
// Example 1:
// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
//
// Example 2:
// Input: root = [1,2]
// Output: 1

package easy

func diameterOfBinaryTree(root *TreeNode) int {
	depths := make(map[*TreeNode]int)

	diameter := 0

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	visited := make([]bool, 0)
	visited = append(visited, false)

	var left int
	var right int

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		v := visited[len(visited)-1]
		visited = visited[0 : len(visited)-1]

		if node == nil {
			continue
		}

		if v == true {
			left = 0
			if node.Left != nil {
				left = depths[node.Left]
			}

			right = 0
			if node.Right != nil {
				right = depths[node.Right]
			}

			max := left
			if right > max {
				max = right
			}

			depths[node] = max + 1

			if (left + right) > diameter {
				diameter = left + right
			}

		} else {
			stack = append(stack, node, node.Left, node.Right)
			visited = append(visited, true, false, false)
		}
	}

	return diameter
}

// Another post order approach
// func diameterOfBinaryTree(root *TreeNode) int {
//     stack := make([]*TreeNode, 0)
//     curr := root

//     nodeToHeight := make(map[*TreeNode]int)

//     diameter := 0

//     var lastVisited *TreeNode
//     for (len(stack) != 0) || (curr != nil) {
//         if curr != nil {
//             stack = append(stack, curr)
//             curr = curr.Left
//         } else {
//             x := stack[len(stack)-1]
//             if (x.Right != nil) && (x.Right != lastVisited) {
//                 curr = x.Right
//             } else {
//                 var left int = 0
//                 if x.Left != nil {
//                     left = nodeToHeight[x.Left]
//                 }

//                 var right int = 0
//                 if x.Right != nil {
//                     right = nodeToHeight[x.Right]
//                 }

//                 var maxLenght int = left
//                 if right > maxLenght {
//                     maxLenght = right
//                 }
//                 nodeToHeight[x] = maxLenght + 1

//                 if (left + right) > diameter {
//                     diameter = left + right
//                 }

//                 lastVisited = x
//                 stack = stack[0:len(stack)-1]
//             }
//         }
//     }

//     return diameter
// }

func diameterOfBinaryTreeRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	recursiveFunc(root, &max)

	return max
}

func recursiveFunc(root *TreeNode, max *int) (depth int) {
	if root == nil {
		return 0
	}

	left := recursiveFunc(root.Left, max)
	right := recursiveFunc(root.Right, max)

	if *max < left+right {
		*max = left + right
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
