package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
1669. 合并两个链表
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。

下图中蓝色边和节点展示了操作后的结果：


请你返回结果链表的头指针。



示例 1：



输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]
解释：我们删除 list1 中第三和第四个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。
示例 2：


输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
解释：上图中蓝色的边和节点为答案链表。


提示：

3 <= list1.length <= 104
1 <= a <= b < list1.length - 1
1 <= list2.length <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-in-between-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cur := list1
	for i := 0; i < a-1; i++ {
		cur = cur.Next
	}
	first := cur
	for i := 0; i < b-a+1; i++ {
		cur = cur.Next
	}
	second := list2
	for second != nil && second.Next != nil {
		second = second.Next
	}
	second.Next = cur.Next
	first.Next = list2
	return list1
}

func TestmergeInBetween() {
	tests := []struct {
		list1 []int
		a     int
		b     int
		list2 []int
		res   []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5},
			3, 4,
			[]int{1000000, 1000001, 1000002},
			[]int{0, 1, 2, 1000000, 1000001, 1000002, 5},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			1, 5,
			[]int{1000000, 1000001, 1000002},
			[]int{0, 1000000, 1000001, 1000002},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			3, 3,
			[]int{1000000, 1000001, 1000002},
			[]int{0, 1, 2, 1000000, 1000001, 1000002, 4, 5},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			2, 5,
			[]int{1000000, 1000001, 1000002, 1000003, 1000004},
			[]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6},
		},
	}

	for _, test := range tests {
		list1 := CreateList(test.list1)
		list2 := CreateList(test.list2)
		list := mergeInBetween(list1, test.a, test.b, list2)
		res := listToArray(list)
		if common.IntEqual(res, test.res) {
			fmt.Println(true)
		} else {
			fmt.Println(res, test.res)
		}

	}
}
