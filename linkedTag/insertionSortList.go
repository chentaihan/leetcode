package linkedTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
147. 对链表进行插入排序
对链表进行插入排序。

插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
 

示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insertion-sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := head
	tail := head
	for cur := root.Next; cur != nil; {
		next := cur.Next
		if cur.Val >= tail.Val {
			tail.Next = cur
			tail = tail.Next
		} else if cur.Val <= root.Val {
			cur.Next = root
			root = cur
		} else {
			for node := root; node != nil; node = node.Next {
				if cur.Val <= node.Next.Val {
					cur.Next = node.Next
					node.Next = cur
					break
				}
			}
		}
		cur = next
	}
	tail.Next = nil
	return root
}

func TestinsertionSortList() {
	tests := []struct {
		nums []int
		res  []int
	}{
		{
			[]int{1, 3, 5, 7, 2, 4, 6, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{9, 7, 5, 3, 1, 8, 6, 4, 2},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 1, 1, 1, 1, 1,},
			[]int{1, 1, 1, 1, 1, 1,},
		},
		{
			[]int{1, 5, 2, 9, 8, 4, 7, 3, 6},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, test := range tests {
		head := CreateList(test.nums)
		head = insertionSortList(head)
		res := listToArray(head)
		fmt.Println(common.IntEqual(res, test.res))
	}
}
