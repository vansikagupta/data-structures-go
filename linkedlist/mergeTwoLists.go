/*
leetcode #21: https://leetcode.com/problems/merge-two-sorted-lists/

// Inplace and iterative merge implementation
// 1. first check both lists are non-empty; if not, return the non-nil list
// 2. result list = min of both list's heads
// 3. iterate through all elements in both lists till both are non-empty,
// 4. append min node to result list and move forward
// 5. repeat 4 and 5
// 6. append the non-nil remaning list to result list
*/

package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// both lists are non-empty
	// head of resultant list; inplace merge
	head := l2
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		l2 = l2.Next
	}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	// either one or both lists are empty, append the remaining part
	if l1 == nil {
		temp.Next = l2
	} else {
		temp.Next = l1
	}
	return head
}
