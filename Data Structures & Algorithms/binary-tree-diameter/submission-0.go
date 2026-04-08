/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    d , _ := hDiameter(root)

	return d
}

func hDiameter(root *TreeNode) (int , int ){
	var d int

	if root == nil {
		return 0, 0
	}

	dr, hr := hDiameter(root.Right)
	dl, hl := hDiameter(root.Left)

	d = max(dr, dl)


	return max(d , hl + hr ), max(hl, hr) + 1 
}
