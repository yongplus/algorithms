package 合并K个升序链表 /**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a != nil {
			return a
		} else {
			return b
		}
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
func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	i := 0
	for i < len(lists) {
		i++
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}
