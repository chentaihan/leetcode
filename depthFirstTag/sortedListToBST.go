package depthFirstTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/treeTag"
)

//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
// 示例:
//
// 给定的有序链表： [-10, -3, 0, 5, 9],
//
//一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
//
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBSTEx(head *ListNode) *treeTag.TreeNode {
	if head == nil {
		return nil
	}
	array := listToArray(head)
	var build func(array []int) *treeTag.TreeNode
	build = func(array []int) *treeTag.TreeNode {
		if len(array) == 0 {
			return nil
		}
		mid := len(array) / 2
		return &treeTag.TreeNode{
			Val:   array[mid],
			Left:  build(array[:mid]),
			Right: build(array[mid+1:]),
		}
	}

	return build(array)
}

func sortedListToBST(head *ListNode) *treeTag.TreeNode {
	array := listToArray(head)
	if len(array) == 0 {
		return nil
	}
	root := &treeTag.TreeNode{
		Val: array[len(array)/2],
	}
	createNode(root, array[:len(array)/2], array[len(array)/2+1:])
	return root
}

func createNode(root *treeTag.TreeNode, leftArray, rightArray []int) *treeTag.TreeNode {
	if root == nil {
		return nil
	}
	if len(leftArray) > 0 {
		index := len(leftArray) / 2
		root.Left = &treeTag.TreeNode{
			Val: leftArray[index],
		}
		createNode(root.Left, leftArray[:index], leftArray[index+1:])
	}
	if len(rightArray) > 0 {
		index := len(rightArray) / 2
		root.Right = &treeTag.TreeNode{
			Val: rightArray[index],
		}
		createNode(root.Right, rightArray[:index], rightArray[index+1:])
	}
	return nil
}

func listToArray(head *ListNode) []int {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return list
}

func arrayToList(array []int) *ListNode {
	var head, cur *ListNode
	if len(array) == 0 {
		return head
	}
	head = &ListNode{
		Val: array[0],
	}
	cur = head
	for i := 1; i < len(array); i++ {
		cur.Next = &ListNode{
			Val: array[i],
		}
		cur = cur.Next
	}
	return head
}

func sortedListToBSTXX(nums []int) *treeTag.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &treeTag.TreeNode{
		Val:   nums[mid],
		Left:  sortedListToBSTXX(nums[:mid]),
		Right: sortedListToBSTXX(nums[mid+1:]),
	}
}

func TestsortedListToBST() {
	tests := []struct {
		list []int
	}{
		{
			[]int{1},
		},
		{
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3,},
		},
		{
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, test := range tests {
		//head := arrayToList(test.list)
		root := sortedListToBSTXX(test.list)
		array := treeTag.TreeFloor(root)
		fmt.Println(array)
	}
}
