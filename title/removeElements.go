package title

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
删除链表中等于给定值 val 的所有节点。
示例:
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	tmpNode := head
	if tmpNode == nil {
		return nil
	}

	for tmpNode.Next != nil {
		if tmpNode.Next.Val == val {
			tmpNode.Next = tmpNode.Next.Next
		} else {
			tmpNode = tmpNode.Next
		}
	}

	return head
}

func createListNode(val string) *ListNode {
	var rootNode, curNode *ListNode
	for _, c := range val {
		tmpNode := &ListNode{
			Val: int(c - '0'),
		}
		if rootNode == nil {
			rootNode, curNode = tmpNode, tmpNode
		} else {
			curNode.Next = tmpNode
			curNode = curNode.Next
		}
	}
	return rootNode
}

func listToString(l *ListNode) string {
	s := ""
	for l != nil {
		s += strconv.Itoa(l.Val)
		l = l.Next
	}
	return s
}

func TestRemoveElements() {
	tests := []struct {
		sour  string
		dest  string
		rmVal int
	}{
		{
			"11",
			"",
			1,
		},
		{
			"111111111111111111111111111",
			"",
			1,
		},
		{
			"1263456",
			"12345",
			6,
		},
		{
			"",
			"",
			6,
		},
		{
			"1",
			"1",
			6,
		},
		{
			"1",
			"",
			1,
		},
		{
			"123",
			"23",
			1,
		},
		{
			"123",
			"13",
			2,
		},
		{
			"123",
			"12",
			3,
		},
		{
			"12345678912345678901",
			"23456789234567890",
			1,
		},
	}
	for _, test := range tests {
		list := createListNode(test.sour)
		listNode := removeElements(list, test.rmVal)
		fmt.Println(listToString(listNode) == test.dest)
	}

}
