/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if root == nil || p == nil || q == nil {
		return nil
	}
	var node *TreeNode

	if min(p.Val, q.Val) > root.Val {
		node = 	lowestCommonAncestor(root.Right, p, q)
	} else if max(p.Val, q.Val) < root.Val {
		node = lowestCommonAncestor(root.Left, p, q)
	} else {
		node = root
	}

	return node
}
