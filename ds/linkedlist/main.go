package main

import "fmt"

//Node represents each element of the LL
type Node struct {
	value int   // holds data
	next  *Node // reference to next Node in the linked list
}

// LinkedList represents the LL, holds reference to the head node
type LinkedList struct {
	head *Node // reference to head of the linked list
	len  int   // optional; num of nodes in the ll
}

// NewLinkedList creates an instance of LinkedList
func NewLinkedList() LinkedList {
	// returns empty list with len 0
	return LinkedList{}
}

// ***** Common Operations *****

// Append inserts a new node at the end of the linked list
func (l *LinkedList) Append(val int) {
	newNode := Node{value: val}
	// if the LL is empty, set new node as head node
	if l.head == nil {
		l.head = &newNode
		l.len++
		return
	}
	// otherwise, traverse till the last node
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	// 1->2->3->4->nil
	// loop will stop when temp is at 4
	temp.next = &newNode
	l.len++
}

// InsertAt inserts new node at the nth position
func (l *LinkedList) InsertAt(n int, val int) {
	// 0->1->2->3->4->nil
	newNode := Node{value: val}
	//invalid position
	if n < 0 || n > l.len {
		return
	}
	if n == 0 {
		newNode.next = l.head
		l.head = &newNode
		l.len++
		return
	}
	// get at (n-1) th node
	prev := l.head
	for i := 1; i < n; i++ {
		prev = prev.next
	}
	newNode.next = prev.next
	prev.next = &newNode
	l.len++
}

// Print traverses and prints the LL
func (l LinkedList) Print() {
	if l.head == nil || l.len == 0 {
		fmt.Println("No nodes in list")
		return
	}
	for temp := l.head; temp != nil; {
		fmt.Print(temp.value, "->")
		temp = temp.next
	}
	fmt.Println()
}

// Search for a value in Linkedlist and retun position
func (l LinkedList) Search(val int) int {
	/* not required explicitly; get's handled in the for loop
	if l.head == nil {
		fmt.Println("No nodes in list")
		return -1
	}*/
	temp := l.head
	for i := 0; temp != nil; i++ {
		if temp.value == val {
			fmt.Println("Value found at pos ", i)
			return i
		}
		temp = temp.next
	}
	fmt.Println("Value not found in list")
	return -1
}

// DeleteAt removes the nth node from ll
func (l *LinkedList) DeleteAt(n int) {
	// invalid position
	if n < 0 || n >= l.len {
		return
	}
	// if n = 0, remove head node
	if n == 0 {
		l.head = l.head.next
		l.len--
		return
	}
	// find n-1 th node
	prev := l.head
	for i := 1; i < n; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	l.len--
}

// DeleteVal removes the first node with value val
func (l *LinkedList) DeleteVal(val int) {
	if l.head.value == val {
		l.head = l.head.next
		l.len--
		return
	}
	prev := l.head
	for prev.next != nil && prev.next.value != val {
		prev = prev.next
	}
	if prev.next == nil {
		fmt.Println("Value not found in LL")
		return
	}
	prev.next = prev.next.next
	l.len--
}

func main() {
	myLinkedList := NewLinkedList()
	myLinkedList.Append(10)
	myLinkedList.Append(40)
	myLinkedList.Print()
	myLinkedList.InsertAt(1, 20)
	myLinkedList.InsertAt(2, 30)
	myLinkedList.InsertAt(3, 50)
	myLinkedList.Print()
	myLinkedList.DeleteVal(50)
	myLinkedList.DeleteVal(40)
	myLinkedList.Print()
	fmt.Println(myLinkedList.Search(50))
	myLinkedList.DeleteAt(0)
	myLinkedList.Print()
	myLinkedList.DeleteAt(1)
	myLinkedList.Print()
	myLinkedList.DeleteAt(0)
	myLinkedList.Print()

	/*
		***********
		Expected output:
		10->40->
		10->20->30->50->40->
		10->20->30->
		Value not found in list
		-1
		20->30->
		20->
		No nodes in list
		***********
	*/
}
