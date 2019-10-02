package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
82. 删除排序链表中的重复元素 II (done)
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var prev, ret *ListNode
	curValue := cur.Val
	if cur.Val != cur.Next.Val {
		ret, prev = cur, cur
	}
	cur = cur.Next
	for ; cur != nil; cur = cur.Next {
		if cur.Val == curValue {
			continue
		}

		if cur.Next == nil || cur.Val != cur.Next.Val {
			if prev == nil {
				prev = cur
				ret = cur
			} else {
				prev.Next = cur
				prev = prev.Next
			}
		}
		curValue = cur.Val
	}
	if prev != nil {
		prev.Next = nil
	}

	return ret
}

func TestdeleteDuplicates() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1, 1},
			[]int{},
		},
		{
			[]int{1, 2, 2, 3, 3, 4, 5},
			[]int{1, 4, 5},
		},
		{
			[]int{1, 1, 2},
			[]int{2},
		},
		{
			[]int{1, 1, 2, 2, 3},
			[]int{3},
		},
		{
			[]int{1, 1, 2, 3},
			[]int{2, 3},
		},
		{
			[]int{1, 1, 2, 3, 4},
			[]int{2, 3, 4},
		},
		{
			[]int{1, 1, 2, 3, 3},
			[]int{2},
		},
		{
			[]int{1, 1, 2, 3, 4, 4},
			[]int{2, 3},
		},
		{
			[]int{1, 1, 2, 2, 3, 3, 4, 5},
			[]int{4, 5},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5},
			[]int{2, 4, 5},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 6, 6},
			[]int{2, 4, 5},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 5, 6, 7},
			[]int{2, 4, 6, 7},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 7},
			[]int{2, 4, 6},
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		result := deleteDuplicates(list)
		array := listToArray(result)
		fmt.Println(common.IntEqual(test.result, array))
	}
}
