package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
TODO
*/

func dreverseBetween(head *ListNode, m int, n int) *ListNode {
	var prev, end *ListNode
	index := 0
	cur := head
	for ; cur != nil; cur = cur.Next {
		index++
		if index == m {
			break
		}
		prev = cur
	}
	end = cur

	var pos = cur
	for cur != nil && m <= n {
		next := cur.Next
		cur.Next = pos
		pos = cur
		cur = next
		m++
	}

	if prev != nil {
		prev.Next = pos
	} else {
		head = pos
	}

	end.Next = cur
	return head
}

func TestreverseBetween() {
	tests := []struct {
		nums   []int
		result []int
		m, n   int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 4, 3, 2, 5},
			2, 4,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 2, 1, 5},
			1, 4,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			1, 5,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			1, 1,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			2, 2,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			5, 5,
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		result := dreverseBetween(list, test.m, test.n)
		array := listToArray(result)
		isOk := common.IntEqual(test.result, array)

		fmt.Printf("%+v", array)
		fmt.Println("      ", isOk)
	}
}
