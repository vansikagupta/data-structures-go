/*******
leetcode #19: https://leetcode.com/problems/remove-nth-node-from-end-of-list/

0->1->2->3->4->5->nil
Sol 1: find len in first traversal,
       delete (len-n)th (0 index) element in second traversal

Sol 2: Using two pointers, first pointer is at head and second ptr = (head + n)th node
	   if ptr i nil, remove head node, else
       ptr increments by 1
	   keep increasing both pointers by 1 node till ptr = nil,
       at this point first pointer is previous to the element to be deleted

********/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	for i := 0; i < n; i++ {
		ptr = ptr.Next
	}

	// which means element to remove is the head node
	if ptr == nil {
		head = head.Next // there's no previous node
		return head
	}

	prev := head
	ptr = ptr.Next
	for ptr != nil {
		prev = prev.Next
		ptr = ptr.Next
	}
	prev.Next = prev.Next.Next // x gets deleted
	return head
}
