/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var h , t *ListNode 

	for list1 != nil && list2 != nil {
		var tmp *ListNode
		if list1.Val < list2.Val {
			tmp = list1
			list1 = list1.Next
		} else {
			tmp = list2
			list2 = list2.Next
		}
		if h == nil {
			h, t  = tmp, tmp
		}
		t.Next = tmp
		t = t.Next
		t.Next = nil
	}

	if list1 == nil {
		if h == nil {
			h, t = list2, list2
			return h
		}
		t.Next = list2
	}

	if list2 == nil {
		if h == nil {
			h , t = list1, list1
			return h
		}
		t.Next = list1
	}

	return h
}
