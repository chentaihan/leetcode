package linkedTag

/**
234. 回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPalindrome(head *ListNode) bool {
	return _isPalindrome(head, head)
}

func _isPalindrome(head1, head2 *ListNode) bool {
	if head1 != nil {
		result := _isPalindrome(head1.Next, head2) && head1.Val == head2.Val
		if head2 != nil && head2.Next != nil {
			*head2 = *head2.Next
		}
		return result
	}
	return true
}
