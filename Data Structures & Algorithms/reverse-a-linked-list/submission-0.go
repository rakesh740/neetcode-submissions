/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    
	var n, t *ListNode = nil, head

	for t != nil {
		tmp := t.Next
		t.Next = n 
		n = t
		t = tmp
	}

	return n
}
