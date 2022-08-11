package array

type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func NewList[T comparable](values ...T) *ListNode[T] {
	dummyHead := &ListNode[T]{}
	cur := dummyHead
	for _, value := range values {
		node := &ListNode[T]{Val: value}
		cur.Next = node
		cur = cur.Next
	}
	return dummyHead.Next
}
