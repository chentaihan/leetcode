package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
143. 重排链表
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.

示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

分析：
先遍历一次，将链表分成两个链表，一个是正序一个是倒叙
在将两个链表依次合并成一个新链表
*/

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	first := head
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		first = slow
		slow = slow.Next
	}
	first.Next = nil
	first = head
	second := slow
	slow = slow.Next
	second.Next = nil
	for slow != nil {
		next := slow.Next
		slow.Next = second
		second = slow
		slow = next
	}

	for second != nil {
		next := first.Next
		first.Next = second
		first = first.Next
		second = second.Next
		if next == nil {
			break
		}
		first.Next = next
		first = first.Next
	}
}

func TestreorderList() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 4, 2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 5, 2, 4, 3},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 6, 2, 5, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 7, 2, 6, 3, 5, 4},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{1, 8, 2, 7, 3, 6, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 9, 2, 8, 3, 7, 4, 6, 5},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6},
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		reorderList(list)
		array := listToArray(list)
		isOk := common.IntEqual(test.result, array)

		fmt.Printf("%+v", array)
		fmt.Println("      ", isOk)
	}
}
