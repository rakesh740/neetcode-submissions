/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var b, c int

	res := &ListNode{}
	curr := res
	for l1 != nil || l2 != nil || c > 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + c 
		b, c = sum%10, sum/10

		d := &ListNode{Val: b}
		curr.Next = d 
		curr = curr.Next
	}

	return res.Next   
}
