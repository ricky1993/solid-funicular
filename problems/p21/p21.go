package p21

import "github.com/ricky1993/solid-funicular/structures/array"

func mergeTwoLists(list1 *array.ListNode[int], list2 *array.ListNode[int]) *array.ListNode[int] {
	dummyHead := &array.ListNode[int]{}
	cur := dummyHead
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	return dummyHead.Next
}
