/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {

	var res *ListNode

	for i:= 0; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}

	return res
}


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var head, tail *ListNode

	for list1 != nil && list2 != nil {
		var tmp *ListNode

		if list1.Val > list2.Val {
			tmp = list2
			list2 = list2.Next
		} else {
			tmp = list1
			list1 = list1.Next
		}

		if head == nil {
			head = tmp
			tail = tmp
		}
		tail.Next = tmp
		tail = tail.Next
		tail.Next = nil
	}

	if list1 != nil {
		if head == nil {
			head = list1
		} else {
			tail.Next = list1
		}
	}

	if list2 != nil {
		if head == nil {
			head = list2
		} else {
			tail.Next = list2
		}

	}

	return head
}