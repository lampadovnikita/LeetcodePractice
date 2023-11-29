// https://leetcode.com/problems/subtree-of-another-tree/
// 572. Subtree of Another Tree
//
// Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.
//
// Example 1:
// Input: root = [3,4,5,1,2], subRoot = [4,1,2]
// Output: true
//
// Example 2:
// Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
// Output: false

package easy

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type merkleTreeNode struct {
	hash  []byte
	left  *merkleTreeNode
	right *merkleTreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	merkleRoot := buildMerkleTree(root)
	merkleSubRoot := buildMerkleTree(subRoot)

	queue := make([]*merkleTreeNode, 0)
	queue = append(queue, merkleRoot)

	for len(queue) != 0 {
		currNode := queue[0]
		queue = queue[1:]

		if bytes.Compare(currNode.hash, merkleSubRoot.hash) == 0 {
			return true
		}

		if currNode.left != nil {
			queue = append(queue, currNode.left)
		}
		if currNode.right != nil {
			queue = append(queue, currNode.right)
		}
	}

	return false
}

func buildMerkleTree(root *TreeNode) (merkleRoot *merkleTreeNode) {
	if root == nil {
		return nil
	}

	merkleRoot = &merkleTreeNode{
		left:  buildMerkleTree(root.Left),
		right: buildMerkleTree(root.Right),
	}

	var leftBytes []byte
	if merkleRoot.left != nil {
		leftBytes = merkleRoot.left.hash
	}

	var rightBytes []byte
	if merkleRoot.right != nil {
		rightBytes = merkleRoot.right.hash
	}

	valBytes := []byte(strconv.Itoa(root.Val))
	valToHash := append(leftBytes, append(valBytes, rightBytes...)...)

	hash := sha256.Sum256(valToHash)
	merkleRoot.hash = hash[:]

	return merkleRoot
}

func isSubtreeRecursive(root *TreeNode, subRoot *TreeNode) bool {
	if isSame(root, subRoot) == true {
		return true
	}

	if root == nil {
		return false
	}

	return isSubtreeRecursive(root.Left, subRoot) || isSubtreeRecursive(root.Right, subRoot)

}

func isSame(f *TreeNode, s *TreeNode) bool {
	if (f == nil) && (s == nil) {
		return true
	} else if (f == nil) || (s == nil) {
		return false
	}

	if f.Val != s.Val {
		return false
	}

	return isSame(f.Left, s.Left) && isSame(f.Right, s.Right)
}
