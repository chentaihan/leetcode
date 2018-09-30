package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
21. 合并两个有序链表
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var cur, root *ListNode
	if l1.Val < l2.Val {
		cur = l1
		l1 = l1.Next
	} else {
		cur = l2
		l2 = l2.Next
	}
	root = cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	return root
}

func TestmergeTwoLists() {
	tests := []struct {
		l1     []int
		l2     []int
		result []int
	}{
		{
			[]int{1},
			[]int{1},
			[]int{1, 1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
			[]int{1, 1, 2, 2},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
			[]int{1, 1, 2, 2, 3, 3, 4, 4},
		},
		{
			[]int{1, 2, 3,},
			[]int{1, 2, 3, 4},
			[]int{1, 1, 2, 2, 3, 3, 4},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3},
			[]int{1, 1, 2, 2, 3, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3},
			[]int{1, 1, 2, 2, 3, 3, 4, 5},
		},
		{
			[]int{1, 2, 3,},
			[]int{1, 2, 3, 4, 5},
			[]int{1, 1, 2, 2, 3, 3, 4, 5},
		},
	}
	for _, test := range tests {
		l1 := createList(test.l1)
		l2 := createList(test.l2)
		result := mergeTwoLists(l1, l2)
		ret := nodeListToArray(result)
		fmt.Println(common.IntEqual(ret, test.result))
	}
}
