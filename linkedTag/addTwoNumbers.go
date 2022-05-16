package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var buf []int
	node1, node2 := l1, l2
	for node1 != nil && node2 != nil {
		buf = append(buf, node1.Val+node2.Val)
		node1 = node1.Next
		node2 = node2.Next
	}

	var result, pre *ListNode
	if node1 != nil {
		result, node1 = l1, l1
	} else {
		result, node1 = l2, l2
	}
	add := 0
	for i := 0; i < len(buf); i++ {
		val := buf[i] + add
		if val >= 10 {
			node1.Val = val - 10
			add = 1
		} else {
			node1.Val = val
			add = 0
		}
		pre = node1
		node1 = node1.Next
	}
	for pre.Next != nil {
		val := pre.Next.Val + add
		if val >= 10 {
			pre.Next.Val = val - 10
			add = 1
		} else {
			pre.Next.Val = val
			add = 0
		}
		pre = pre.Next
	}
	if add == 1 {
		pre.Next = &ListNode{
			Val: add,
		}
	}
	return result
}

func addTwoNumbersEx(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	add := 0
	node1, node2 := l1, l2
	var pre *ListNode
	for node1 != nil && node2 != nil {
		add += node1.Val + node2.Val
		if add >= 10 {
			node1.Val = add - 10
			add = 1
		} else {
			node1.Val = add
			add = 0
		}
		pre = node1
		node1 = node1.Next
		node2 = node2.Next
	}
	if node2 != nil {
		pre.Next = node2
	}
	for pre.Next != nil {
		val := pre.Next.Val + add
		if val >= 10 {
			pre.Next.Val = val - 10
			add = 1
		} else {
			pre.Next.Val = val
			add = 0
		}
		pre = pre.Next
	}
	if add == 1 {
		pre.Next = &ListNode{
			Val: add,
		}
	}
	return l1
}

func TestaddTwoNumbers() {
	tests := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{
		{
			[]int{9, 8},
			[]int{1},
			[]int{0, 9},
		},
		{
			[]int{5},
			[]int{5},
			[]int{0, 1},
		},
		{
			[]int{5, 5, 5, 5, 5, 5},
			[]int{5, 5, 5, 5, 5, 5},
			[]int{0, 1, 1, 1, 1, 1, 1},
		},
		{
			[]int{5, 5, 5, 5, 5, 5},
			[]int{5, 5, 5, 5, 5, 5, 9, 9, 9, 9},
			[]int{0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{2, 4, 6},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			[]int{2, 4, 6, 8, 0, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5, 9},
			[]int{2, 4, 6, 8, 0, 0, 1},
		},
	}

	for _, test := range tests {
		l1 := CreateList(test.nums1)
		l2 := CreateList(test.nums2)
		result := addTwoNumbersEx(l1, l2)
		array := listToArray(result)
		fmt.Println(common.IntEqual(array, test.result))
	}

}
