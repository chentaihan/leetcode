package linkedTag

/**
24. 两两交换链表中的节点(完成)
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/


func swapPairs(head *ListNode) *ListNode {
	next := head
	var first, second,cur *ListNode
	for next != nil {
		first = next
		second = first.Next
		if second == nil {
			 break
		}
		next = second.Next
		second.Next = first
		first.Next = next
		if cur == nil {
			head = second
		}else {
			cur.Next = second
		}
		cur = first
	}
	return head
}
