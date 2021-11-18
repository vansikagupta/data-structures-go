// leetcode #141 :https://leetcode.com/problems/linked-list-cycle/

package main

/*
Two pointer approach; Floyd's Algorithm
1. slow pointer moves one node at a time
2. fast pointer moves two nodes at a time
3. If both pointers meet at any point, cycle is detected
4. else if, we reach end of list, no cycle is present
*/
// Time O(n), Space O(1)

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	// confirms that two steps can be taken without going out of index
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
