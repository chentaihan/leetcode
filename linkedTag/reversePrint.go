package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
面试题06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/

func reversePrint(head *ListNode) []int {
	count := 0
	for node := head; node != nil; node = node.Next {
		count++
	}
	result := make([]int, count)
	for node := head; node != nil; node = node.Next {
		count--
		result[count] = node.Val
	}
	return result
}

func reversePrintEx(head *ListNode) []int {
	result := make([]int, 0)
	var f func(head *ListNode)
	f = func(head *ListNode) {
		if head != nil {
			f(head.Next)
			result = append(result, head.Val)
		}
	}
	f(head)
	return result
}

func TestreversePrint() {
	tests := []struct {
		nums   []int
		result []int
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
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		list := CreateList(test.nums)
		result := reversePrintEx(list)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
