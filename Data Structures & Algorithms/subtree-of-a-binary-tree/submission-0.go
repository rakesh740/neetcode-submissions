/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil || root == nil {
		return subRoot == root
	}
	if root.Val == subRoot.Val && isSameTree(root, subRoot) {
		return true
	}
		
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) 
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
    
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}