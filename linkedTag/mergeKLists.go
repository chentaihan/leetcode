package linkedTag

/**
23. 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

分析：两两合并
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeList(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func mergeList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	var root *ListNode
	if left.Val < right.Val {
		root = left
		left = left.Next
	} else {
		root = right
		right = right.Next
	}
	cur := root
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	} else {
		cur.Next = right
	}
	return root
}
