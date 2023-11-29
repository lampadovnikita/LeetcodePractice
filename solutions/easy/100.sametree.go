// https://leetcode.com/problems/same-tree/
// 100. Same Tree
//
// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//
// Example 1:
// Input: p = [1,2,3], q = [1,2,3]
// Output: true
//
// Example 2:
// Input: p = [1,2], q = [1,null,2]
// Output: false
//
// Example 3:
// Input: p = [1,2,1], q = [1,1,2]
// Output: false

package easy

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil) || (q == nil) {
		if p != q {
			return false
		}

		return true
	}

	qp := make([]*TreeNode, 0)
	qq := make([]*TreeNode, 0)

	qp = append(qp, p)
	qq = append(qq, q)

	var currq *TreeNode
	var currp *TreeNode

	for (len(qp) != 0) && (len(qq) != 0) {
		currp = qp[0]
		qp = qp[1:]

		currq = qq[0]
		qq = qq[1:]

		if currq.Val != currp.Val {
			return false
		}

		if currp.Left != nil {
			qp = append(qp, currp.Left)
		}
		if currq.Left != nil {
			qq = append(qq, currq.Left)
		}

		if len(qq) != len(qp) {
			return false
		}

		if currp.Right != nil {
			qp = append(qp, currp.Right)
		}
		if currq.Right != nil {
			qq = append(qq, currq.Right)
		}

		if len(qq) != len(qp) {
			return false
		}
	}

	if (len(qp) != 0) || (len(qq) != 0) {
		return false
	}

	return true
}
