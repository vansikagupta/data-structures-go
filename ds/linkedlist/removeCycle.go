// leetcode #142 : https://leetcode.com/problems/linked-list-cycle-ii/

package main

/* Hash Table approach

1. for each node, check if it is present in the table
2. the first node to be repeated is the beginning of cycle
3. If node, not present, insert and continue

Space: O(n) Time: O(n)
*/

/* Floyd's Algo
1. Using two pointer approach, first detect cycle
2. keep one pointer where slow and fast met, keep one pointer at head
3. move both by one step
4. the meeting point is the cycle beginning

Space: O(1) Time: O(n)
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// cycle detected
			return cycleBeginsAt(head, slow)
		}
	}
	return nil
}

func cycleBeginsAt(head, node *ListNode) *ListNode {
	for head != node {
		head = head.Next
		node = node.Next
	}
	return node
}

/*
Why this works?
See README.md
fast = 2 * slow
m + x*n + k = 2(m + y*n +k)
m = (x-2y)*n - k
m = C*n - k

which essentially means m steps from head and after some complete cycles minus k steps already taken,
 we will meet at starting point of cycle
*/
