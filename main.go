package main

type Node struct {
	Val  int
	Next *Node
}

func LastK(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	fast := head
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	
}
