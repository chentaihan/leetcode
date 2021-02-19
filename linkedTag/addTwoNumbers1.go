package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
445. 两数相加 II
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var sum1, sum2 []int
	for item := l1; item != nil; item = item.Next {
		sum1 = append(sum1, item.Val)
	}
	for item := l2; item != nil; item = item.Next {
		sum2 = append(sum2, item.Val)
	}
	prev := 0
	if len(sum1) < len(sum2) {
		sum1, sum2 = sum2, sum1
		l1, l2 = l2, l1
	}
	index1, index2 := len(sum1)-1, len(sum2)-1
	for index1 >= 0 && index2 >= 0 {
		val := sum1[index1] + sum2[index2] + prev
		if val >= 10 {
			sum1[index1] = val - 10
			prev = 1
		} else {
			sum1[index1] = val
			prev = 0
		}
		index1--
		index2--
	}
	for ; index1 >= 0; index1-- {
		val := sum1[index1] + prev
		if val >= 10 {
			sum1[index1] = val - 10
			prev = 1
		} else {
			sum1[index1] = val
			prev = 0
		}
	}
	index := 0
	for item := l1; item != nil; item = item.Next {
		item.Val = sum1[index]
		index++
	}
	if prev > 0 {
		return &ListNode{Val: prev, Next: l1}
	}
	return l1
}

//输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7

func TestaddTwoNumbers1() {
	tests := []struct {
		num1 []int
		num2 []int
		res  []int
	}{
		{
			[]int{1},
			[]int{9, 9},
			[]int{1, 0, 0},
		},
		{
			[]int{0},
			[]int{7, 2},
			[]int{7, 2},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{
			[]int{7, 2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 8, 0, 7},
		},
		{
			[]int{5, 5, 5},
			[]int{4, 4, 5},
			[]int{1, 0, 0, 0},
		},
		{
			[]int{5, 5, 5},
			[]int{5, 5, 5},
			[]int{1, 1, 1, 0},
		},
		{
			[]int{5, 5, 5},
			[]int{5, 5, 6},
			[]int{1, 1, 1, 1},
		},
		{
			[]int{5, 5, 5},
			[]int{0},
			[]int{5, 5, 5},
		},
	}
	for index, test := range tests {
		list1 := CreateList(test.num1)
		list2 := CreateList(test.num2)
		res := addTwoNumbers1(list1, list2)
		arr := listToArray(res)
		if !common.IntEqual(arr, test.res) {
			fmt.Println(index, "error")
		}
	}
}
