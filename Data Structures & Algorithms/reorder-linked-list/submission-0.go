/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	 fast, slow := head.Next, head

   // get to mid
	for fast != nil && fast.Next != nil  {
		fast = fast.Next.Next
		slow = slow.Next
	}
	sec := slow.Next
	slow.Next = nil

	// reverse last part 
	var prev *ListNode
	for sec != nil {
		tmp := sec.Next
		sec.Next = prev
		prev = sec
		sec = tmp
	}

	// merge both list alternatively
	curr := head
	for curr != nil && prev != nil {
		tmp1, tmp2 := curr.Next, prev.Next
		
		curr.Next = prev
		prev.Next = tmp1


		curr, prev = tmp1, tmp2
	}
	
}
