package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
897. 递增顺序查找树
*/

func increasingBST(root *TreeNode) *TreeNode {
	if root != nil {
		result := &TreeNode{}
		_increasingBST(root, result)
		return result.Right
	}
	return nil
}

func _increasingBST(root *TreeNode, result *TreeNode) {
	if root != nil {
		_increasingBST(root.Left, result)
		for result.Right != nil {
			result = result.Right
		}
		result.Right = &TreeNode{Val: root.Val}
		_increasingBST(root.Right, result)
	}
}

func increasingBSTEx(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	result := &TreeNode{}
	list := result
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		list.Right = &TreeNode{Val: cur.Val}
		list = list.Right
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return result.Right
}

func TestincreasingBST() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			[]int{2, 1, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{5, 3, 6, 2, 4, 7},
			[]int{2, 3, 4, 5, 6, 7},
		},
		{
			[]int{5, 3, 7, 2, 4, 6, 8},
			[]int{2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{5, 3, 7, 2, 4, 6, 8, 1, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		root = increasingBSTEx(root)
		result := TreeFloor(root)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
