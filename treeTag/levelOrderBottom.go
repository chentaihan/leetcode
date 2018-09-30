package treeTag

/**
107. 二叉树的层次遍历 II
 */

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

func levelOrderBottom(root *TreeNode) [][]int {
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
	start, end := 0, len(ret)-1
	for start < end {
		ret[start], ret[end] = ret[end], ret[start]
		start++
		end--
	}
	return ret
}

func TestlevelOrderBottom() {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{},
			[][]int{{}},
		},
		{
			[]int{7},
			[][]int{{7}},
		},
		{
			[]int{7, 3},
			[][]int{{3}, {7}},
		},
		{
			[]int{7, 15},
			[][]int{{15}, {7}},
		},
		{
			[]int{7, 3, 15},
			[][]int{{3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 9},
			[][]int{{9}, {3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 20},
			[][]int{{20}, {3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 9, 20},
			[][]int{{9, 20}, {3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1},
			[][]int{{1, 9, 20}, {3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4},
			[][]int{{1, 4, 9, 20}, {3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4},
			[][]int{{1, 4, 9, 20}, {3, 15}, {7}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4, 30},
			[][]int{{30}, {1, 4, 9, 20}, {3, 15}, {7}},
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		ret := levelOrderBottom(root)
		for index, list := range ret {
			fmt.Println(common.IntEqual(test.result[index], list))
		}
	}
}
