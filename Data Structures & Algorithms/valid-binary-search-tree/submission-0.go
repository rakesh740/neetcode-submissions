/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(node *TreeNode, left , right int)  bool {
	
	if node == nil {
		return true
	}

	val := node.Val

	if val <= left || val >= right {
		return false
	}

	return isValid(node.Left, left, val) && isValid(node.Right, val, right)
}