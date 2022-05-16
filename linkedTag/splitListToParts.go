package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
725. 分隔链表
给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。

每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。

这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。

返回一个符合上述规则的链表的列表。

举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]

示例 1：

输入:
root = [1, 2, 3], k = 5
输出: [[1],[2],[3],[],[]]
解释:
输入输出各部分都应该是链表，而不是数组。
例如, 输入的结点 root 的 val= 1, root.next.val = 2, \root.next.next.val = 3, 且 root.next.next.next = null。
第一个输出 output[0] 是 output[0].val = 1, output[0].next = null。
最后一个元素 output[4] 为 null, 它代表了最后一个部分为空链表。
示例 2：

输入:
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
解释:
输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。


提示:

root 的长度范围： [0, 1000].
输入的每个节点的大小范围：[0, 999].
k 的取值范围： [1, 50].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-linked-list-in-parts
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func splitListToParts(root *ListNode, k int) []*ListNode {
	result := make([]*ListNode, k)
	if root == nil {
		return result
	}
	l := 0
	cur := root
	for cur != nil {
		l++
		cur = cur.Next
	}

	cur = root
	left := l % k
	count := l / k
	prev := cur
	for index := 0; index < k; index++ {
		result[index] = cur
		for i := count - 1; i >= 0; i-- {
			prev = cur
			cur = cur.Next
		}
		if index < left {
			prev = cur
			cur = cur.Next
		}
		prev.Next = nil
		if cur == nil {
			break
		}
	}
	return result
}

func TestsplitListToParts() {
	tests := []struct {
		nums   []int
		result [][]int
		k      int
	}{
		{
			[]int{1},
			[][]int{{1}},
			1,
		},
		{
			[]int{1, 2},
			[][]int{{1}, {2}},
			2,
		},
		{
			[]int{1, 2, 3},
			[][]int{{1}, {2}, {3}},
			3,
		},
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}},
			1,
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{{1}, {2}, {3}, {4}},
			4,
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{{1, 2}, {3, 4}},
			2,
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{{1}, {2}, {3}, {4}, {}, {}},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}, {10}},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}, {9, 10}},
			4,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			2,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			1,
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		result := splitListToParts(list, test.k)
		for index, ret := range result {
			array := listToArray(ret)
			isOk := common.IntEqual(test.result[index], array)

			fmt.Printf("%+v   ,k=%d", array, test.k)
			fmt.Println("      ", isOk)
		}

	}
}
