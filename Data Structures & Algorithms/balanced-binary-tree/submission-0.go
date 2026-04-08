/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
   _, b := disBalanced(root)
   return b
}

func disBalanced(root *TreeNode) (int, bool) {
	 if root == nil {
        return 0, true
    }
	dR, bR := disBalanced(root.Right)
		dL, bL := disBalanced(root.Left)

    return max(dR, dL) + 1, bR && bL && math.Abs(float64(dR - dL)) <= 1.00
}