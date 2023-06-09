package treeTag

/**
102. 二叉树的层次遍历
*/

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := QueueEx{}
	ret := [][]int{}
	queue.Push(root, 0)
	for !queue.Empty() {
		node, level := queue.Pop()
		if node.Left != nil {
			queue.Push(node.Left, level+1)
		}
		if node.Right != nil {
			queue.Push(node.Right, level+1)
		}
		if level >= len(ret) {
			ret = append(ret, []int{node.Val})
		} else {
			ret[level] = append(ret[level], node.Val)
		}
	}
	return ret
}

func levelOrderEx(root *TreeNode) [][]int {
	var list [][]int
	_levelOrderEx(root, &list, 0)
	return list
}

func _levelOrderEx(root *TreeNode, list *[][]int, level int) {
	if root != nil {
		for len(*list) <= level {
			*list = append(*list, []int{})
		}
		(*list)[level] = append((*list)[level], root.Val)
		_levelOrderEx(root.Left, list, level+1)
		_levelOrderEx(root.Right, list, level+1)
	}
}

func levelOrderNew(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue1, queue2 []*TreeNode
	queue1 = append(queue1, root)
	var result [][]int
	var list []int
	for len(queue1) > 0 {
		cur := queue1[0]
		if cur.Left != nil {
			queue2 = append(queue2, cur.Left)
		}
		if cur.Right != nil {
			queue2 = append(queue2, cur.Right)
		}
		list = append(list, cur.Val)
		queue1 = queue1[1:]
		if len(queue1) == 0 {
			result = append(result, list)
			list = nil
			queue1 = queue2
			queue2 = nil
		}
	}
	return result
}

func TestlevelOrder() {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		//{
		//	[]int{},
		//	[][]int{{}},
		//},
		//{
		//	[]int{7},
		//	[][]int{{7}},
		//},
		//{
		//	[]int{7, 3},
		//	[][]int{{7}, {3}},
		//},
		//{
		//	[]int{7, 15},
		//	[][]int{{7}, {15}},
		//},
		{
			[]int{7, 3, 15},
			[][]int{{7}, {3, 15}},
		},
		{
			[]int{7, 3, 15, 9},
			[][]int{{7}, {3, 15}, {9}},
		},
		{
			[]int{7, 3, 15, 20},
			[][]int{{7}, {3, 15}, {20}},
		},
		{
			[]int{7, 3, 15, 9, 20},
			[][]int{{7}, {3, 15}, {9, 20}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1},
			[][]int{{7}, {3, 15}, {9, 20, 1}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4},
			[][]int{{7}, {3, 15}, {9, 20, 1, 4}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4},
			[][]int{{7}, {3, 15}, {9, 20, 1, 4}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4, 30},
			[][]int{{7}, {3, 15}, {9, 20, 1, 4}, {30}},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		ret := levelOrderNew(root)
		for index, list := range ret {
			if common.IntEqual(test.result[index], list) {
				fmt.Println(true)
			} else {
				fmt.Println(test.result[index], list)
			}
		}
	}
}
