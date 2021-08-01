package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
206. 反转链表
反转一个单链表。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode = nil
	cur := head
	next := head.Next
	for cur != nil {
		cur.Next = prev
		if next == nil {
			break
		}
		prev = cur
		cur = next
		next = next.Next
	}
	return cur
}

func reverseListEx(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode = nil
	cur := head
	next := head.Next
	for cur != nil {
		cur.Next, prev = prev, cur
		if next == nil {
			break
		}
		cur, next = next, next.Next
	}
	return cur
}

func TestreverseList() {
	tests := []struct {
		list []int
		ret  []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{6, 5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		head := CreateList(test.list)
		head = reverseListEx(head)
		ret := NodeListToArray(head)
		fmt.Println(common.IntEqual(ret, test.ret))
	}
}
