/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
		k := n - 1
    var prev *ListNode
	 curr :=  head

	for curr.Next != nil {
        if k == 0 {
            prev =  head
        }
		if k < 0 {
			prev = prev.Next
		}
		k--
		curr = curr.Next
	}

	if prev == nil {
		return head.Next
	}

	prev.Next = prev.Next.Next

	return head
}
