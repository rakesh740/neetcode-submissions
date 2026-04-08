/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		 return [][]int{}
	}
	res := [][]int{}


	// create a queue 

	// check q not empty then 
	// add to the q
	// loop through each level 

	q := []*TreeNode{root}
	
    
	for len(q) != 0 {
		qLen := len(q)
		var l []int

		for i:= 0;i < qLen; i++ {
			e := q[0]
			q = q[1:]

			l = append(l, e.Val)

			if e.Left != nil {
				q =append(q, e.Left)
			}
			if e.Right != nil {
				q = append(q, e.Right)
			}
		}
		res = append( res, l)
	}

	return res
}
