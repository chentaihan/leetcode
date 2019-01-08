package linkedTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
206. 反转链表
反转一个单链表。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
 */

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	var prev *common.ListNode = nil
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

func TestreverseList(){
	tests := []struct{
		list []int
		ret []int
	} {
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1,2},
			[]int{2,1},
		},
		{
			[]int{1,2,3},
			[]int{3,2,1},
		},
		{
			[]int{1,2,3,4},
			[]int{4,3,2,1},
		},
		{
			[]int{1,2,3,4,5},
			[]int{5,4,3,2,1},
		},
		{
			[]int{1,2,3,4,5,6},
			[]int{6,5,4,3,2,1},
		},
		{
			[]int{1,2,3,4,5,6,7},
			[]int{7,6,5,4,3,2,1},
		},
		{
			[]int{1,2,3,4,5,6,7,8},
			[]int{8,7,6,5,4,3,2,1},
		},
		{
			[]int{1,2,3,4,5,6,7,8,9},
			[]int{9,8,7,6,5,4,3,2,1},
		},
	}

	for _,test := range tests {
		head := common.CreateList(test.list)
		head = reverseList(head)
		ret := common.NodeListToArray(head)
		fmt.Println(common.IntEqual(ret, test.ret))
	}
}