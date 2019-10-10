package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

*/
func deleteDuplicates_s(head *ListNode) *ListNode {
	for cur := head; cur != nil; {
		if cur.Next == nil {
			break
		}
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func TestdeleteDuplicates_s() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1, 1, 1},
			[]int{1},
		},
		{
			[]int{1, 1},
			[]int{1},
		},
		{
			[]int{1, 2, 2, 3, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 1, 2, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 1, 2, 3, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 2, 3, 4, 4},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 1, 2, 2, 3, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 6, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 5, 6, 7, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		result := deleteDuplicates_s(list)
		array := listToArray(result)
		isOk := common.IntEqual(test.result, array)

		fmt.Printf("%+v", array)
		fmt.Println("      ", isOk)
	}
}
