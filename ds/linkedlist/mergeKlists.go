/*
leetcode #23: https://leetcode.com/problems/merge-k-sorted-lists/

Here we are using Divide and Conquer approach
Algorithm:
1. If k = 0, return; no lists to merge
2. If k = 1, return the only list; no merge needed
3. If k = 2, merge two lists (see ds/linkedlist/mergeTwoLists.go)
4. To merge k lists, merge first k/2 lists into list1 and remaining k/2 lists into list2
5. Now we have two lists to merge, list1 and list2, which can be done easily
*/
package main

func mergeKLists(lists []*ListNode) *ListNode {
	len := len(lists)
	// base conditions
	if len == 0 {
		return nil
	}
	if len == 1 {
		return lists[0]
	}
	if len == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	// divide and conquer, recursively
	list1 := mergeKLists(lists[:len/2])
	list2 := mergeKLists(lists[len/2:])
	return mergeTwoLists(list1, list2)
}
