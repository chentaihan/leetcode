package linkedTag

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
