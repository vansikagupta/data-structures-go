// leetcode #206: https://leetcode.com/problems/reverse-linked-list/
//
package main

/*******************
Recursive Approach
Example : 1->2->3->4->5->nil
* continue recursion until last node is reached
* this is the head node of reversed list
5->nil this is a reversed list
* when at current node, initial next becomes new previous
* set current node.Next to nil
5->4->nil
*/

func reverseList(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node // base condition that gives head of new reversed list
	}
	head := reverseList(node.Next)
	// current next becomes the new prev
	node.Next.Next = node
	node.Next = nil
	return head
}

/******************
Iterative approach

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // node to reverse
    node := head
    var newNext, currNext *ListNode
    // while node to reverse is not nil
    for node != nil {
        currNext = node.Next // store the current next of node before breaking link
        node.Next = newNext // reverse node by assigning new next to node
        newNext = node // update the new next
        node = currNext // change node to reverse
    }
    //head of the reversed list
    return newNext
}
*/
