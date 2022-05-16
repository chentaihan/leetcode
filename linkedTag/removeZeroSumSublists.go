package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
1171. 从链表中删去总和值为零的连续节点
给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。

删除完毕后，请你返回最终结果链表的头节点。



你可以返回任何满足题目要求的答案。

（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

示例 1：

输入：head = [1,2,-3,3,1]
输出：[3,1]
提示：答案 [1,2,1] 也是正确的。
示例 2：

输入：head = [1,2,3,-3,4]
输出：[1,2,4]
示例 3：

输入：head = [1,2,3,-3,-2]
输出：[1]

提示：

给你的链表中可能有 1 到 1000 个节点。
对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：从当前位置往前累加，如当前位置为i，第一个元素为0~i的和,第二个元素为1~i的和，第i个元素为第i个元素的值，
将这些元素加入一个数组Slice中，当某个元素的值为0时，就将这个元素和之后的所有元素从数组中删除
*/

type listRemoveItem struct {
	node  *ListNode
	value int
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	root := head
	for root.Next != nil && root.Val == 0 {
		root = root.Next
	}
	var prev *ListNode
	var list []*listRemoveItem
	for item := head; item != nil; item = item.Next {
		if item.Val == 0 {
			if prev == nil {
				root = item.Next
			} else {
				prev.Next = item.Next
			}
		} else {
			isFound := false
			for i := 0; i < len(list); i++ {
				list[i].value += item.Val
				if list[i].value == 0 {
					if list[i].node == nil {
						root = item.Next
						prev = nil
					} else {
						prev = list[i].node
						list[i].node.Next = item.Next
					}
					list = list[:i]
					isFound = true
					break
				}
			}
			if !isFound {
				list = append(list, &listRemoveItem{
					node:  prev,
					value: item.Val,
				})
				prev = item
			}
		}
	}
	return root
}

func TestremoveZeroSumSublists() {
	tests := []struct {
		list []int
		res  []int
	}{
		{
			[]int{2, 2, -2, 1, -1, -1},
			[]int{2, -1},
		},
		{
			[]int{1, 2, 3, 4, -4, -3, -2, -1},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, -4, -2, -3, -1},
			[]int{},
		},
		{
			[]int{1, 2, -3, 3, 1},
			[]int{3, 1},
		},
		{
			[]int{0, 0, 0},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, -10},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, -10, 5},
			[]int{5},
		},
		{
			[]int{1, 1, 2, 3, 4, -10, 5},
			[]int{1, 5},
		},
		{
			[]int{1, 1, 0, 2, 3, 4, -10, 5},
			[]int{1, 5},
		},
		{
			[]int{3, -3 - 3, 3},
			[]int{},
		},
		{
			[]int{0, 3, -3 - 3, 3, 5},
			[]int{5},
		},
		{
			[]int{0, 3, -3, -5, -3, 3, 5},
			[]int{},
		},
		{
			[]int{0, 3, -3, -5, -3, 3, 5, 1, -1},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -55},
			[]int{},
		},
		{
			[]int{-55, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{},
		},
	}
	for index, test := range tests {
		list := CreateList(test.list)
		res := removeZeroSumSublists(list)
		resList := listToArray(res)
		if !common.IntEqual(test.res, resList) {
			fmt.Println(index, "error")
		}
	}
}
