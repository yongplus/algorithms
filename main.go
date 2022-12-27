package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	head := &ListNode{}
	tail := head
	aPtr := a
	bPtr := b
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr != nil {
		tail.Next = aPtr
	} else {
		tail.Next = bPtr
	}
	return head.Next
}

func merge(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func main() {

}
