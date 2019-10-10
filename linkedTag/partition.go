package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
86. 分隔链表
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

*/

func partition(head *ListNode, x int) *ListNode {
	var left, leftRoot, right, rightRoot *ListNode
	for head != nil {
		if head.Val < x {
			if left == nil {
				left = head
				leftRoot = left
			} else {
				left.Next = head
				left = left.Next
			}
		} else {
			if right == nil {
				right = head
				rightRoot = right
			} else {
				right.Next = head
				right = right.Next
			}
		}
		head = head.Next
	}
	if leftRoot == nil {
		return rightRoot
	}
	if right != nil {
		right.Next = nil
	}

	left.Next = rightRoot
	return leftRoot
}

func Testpartition() {
	tests := []struct {
		nums   []int
		result []int
		x      int
	}{
		{
			[]int{1, 3, 1},
			[]int{1, 1, 3},
			2,
		},
		{
			[]int{4, 1},
			[]int{1, 4},
			2,
		},
		{
			[]int{1, 4},
			[]int{1, 4},
			5,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
			[]int{1, 3, 2, 4, 5, 7, 9, 6, 8, 10},
			5,
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		result := partition(list, test.x)
		array := listToArray(result)
		isOk := common.IntEqual(test.result, array)

		fmt.Printf("%+v", array)
		fmt.Println("      ", isOk)
	}
}
