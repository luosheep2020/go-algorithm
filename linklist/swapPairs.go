package main

func swapPairs(head *ListNode) *ListNode {
	return nil
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs1(next.Next)
	next.Next = head
	return next
}
