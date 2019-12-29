package linkedTag

/*
148. 排序链表
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：链表快排
 */

func sortList(head *ListNode) *ListNode {
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