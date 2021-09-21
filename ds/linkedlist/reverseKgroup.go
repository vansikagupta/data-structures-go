// leetcode #25 : https://leetcode.com/problems/reverse-nodes-in-k-group/

package main

/*
Input : head of list and k(int)
Example: 1->2->3->4->5->6->nil k=3
1. take k steps and reach kth element
	head = 1, kth = 3
2. if this is the last group; kth.Next == nil
	a. if last group has less than k elements, do nothing return head
	b. rest of the list is nil
3. Reverse rest of the list in k groups and store resultant list head
4. set kth.Next = nil and reverse current group
	reverse(head) returns 3->2->1->nil
5. link both current reversed group and resultant reversed list
	set head.Next = reverse rest of the list in k groups

Time: reversing n/k number of groups with k elements
n/k * k
O(n)
Space: recursive calls n/k times, call stack space O(n/k)
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	kth := head
	i := 1
	for ; i < k && kth.Next != nil; i++ {
		kth = kth.Next
	}
	var restOfList *ListNode
	//base condition; we are in the last group
	if kth.Next == nil {
		if i < k {
			//no need to reverse current group
			return head
		}
		restOfList = nil
	} else {
		// reverse remaining list in k groups
		restOfList = reverseKGroup(kth.Next, k)
		// break kth node's link
		kth.Next = nil
	}
	revHead := reverseList(head)

	// link both sublists
	head.Next = restOfList
	return revHead
}
