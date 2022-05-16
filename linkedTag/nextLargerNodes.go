package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
1019. 链表中的下一个更大节点
给出一个以头节点 head 作为第一个节点的链表。链表中的节点分别编号为：node_1, node_2, node_3, ... 。

每个节点都可能有下一个更大值（next larger value）：对于 node_i，如果其 next_larger(node_i) 是 node_j.val，那么就有 j > i 且  node_j.val > node_i.val，而 j 是可能的选项中最小的那个。如果不存在这样的 j，那么下一个更大值为 0 。

返回整数答案数组 answer，其中 answer[i] = next_larger(node_{i+1}) 。

注意：在下面的示例中，诸如 [2,1,5] 这样的输入（不是输出）是链表的序列化表示，其头节点的值为 2，第二个节点值为 1，第三个节点值为 5 。



示例 1：

输入：[2,1,5]
输出：[5,5,0]
示例 2：

输入：[2,7,4,3,5]
输出：[7,0,5,5,0]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-node-in-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func nextLargerNodes_old(head *ListNode) []int {
	if head == nil {
		return nil
	}
	start, cur := head, head.Next
	start.Next = nil
	count := 1
	for cur != nil {
		next := cur.Next
		cur.Next = start
		start = cur
		cur = next
		count++
	}
	result := make([]int, count)
	max := start.Val
	start = start.Next
	index := count - 1
	for start != nil {
		if result[index] > max {
			result[index] = 0
		} else {
			result[index] = max
		}
		if start.Val > max {
			max = start.Val
		}
		start = start.Next
		index--
	}
	return result
}

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var result []int
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	count := len(result) - 1
	max := result[count]
	result[count] = 0
	for index := count - 1; index >= 0; index-- {
		curValue := result[index]
		if result[index] > max {
			result[index] = 0
		} else {
			result[index] = max
		}
		if curValue > max {
			max = curValue
		}
	}
	return result
}

func TestnextLargerNodes() {
	tests := []struct {
		nums   []int
		result []int
	}{
		//{
		//	[]int{1, 1},
		//	[]int{1, 0},
		//},
		{
			[]int{2, 7, 4, 3, 5},
			[]int{7, 0, 5, 5, 0},
		},
		{
			[]int{1, 1, 2},
			[]int{2, 2, 0},
		},
		{
			[]int{1, 1, 2, 2, 3},
			[]int{3, 3, 3, 3, 0},
		},
		{
			[]int{4, 3, 2, 1},
			[]int{0, 0, 0, 0},
		},
		{
			[]int{1, 5, 2, 3, 4},
			[]int{5, 0, 4, 4, 0},
		},
		{
			[]int{7, 7, 1, 2, 3, 5, 4, 4, 3, 2, 1},
			[]int{7, 0, 5, 5, 5, 0, 4, 0, 0, 0, 0},
		},
		{
			[]int{1, 1, 2, 3, 3, 4, 5, 5, 6, 7},
			[]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 0},
		},
		{
			[]int{1, 1, 2, 3, 3, 8, 5, 5, 6, 7},
			[]int{8, 8, 8, 8, 8, 0, 7, 7, 7, 0},
		},
	}

	for _, test := range tests {
		list := CreateList(test.nums)
		result := nextLargerNodes(list)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
