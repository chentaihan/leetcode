package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
94. 二叉树的中序遍历
非递归实现
 */

func inorderTraversal(root *TreeNode) []int {
	list := []int{}
	if root == nil {
		return list
	}
	stack := NewStackTree()
	stack.Push(root)
	for !stack.Empty() {
		//左子树，一步到叶子节点
		for root.Left != nil {
			stack.Push(root.Left)
			root = root.Left
		}
        //从叶子节点往根节点迭代
		for !stack.Empty() {
			//根节点
			root = stack.Pop()
			list = append(list, root.Val)
			//右子树
			if root.Right != nil {
				stack.Push(root.Right)
				root = root.Right
				break
			}
		}
	}
	return list
}

func TestinorderTraversal() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{2, 4},
			[]int{2, 4},
		},
		{
			[]int{4, 2},
			[]int{2, 4},
		},
		{
			[]int{4, 2, 7},
			[]int{2, 4, 7},
		},
		{
			[]int{2, 4, 7},
			[]int{2, 4, 7},
		},
		{
			[]int{4, 2, 7, 1},
			[]int{1, 2, 4, 7},
		},
		{
			[]int{4, 2, 7, 1, 6},
			[]int{1, 2, 4, 6, 7},
		},
		{
			[]int{4, 2, 7, 1, 6, 3},
			[]int{1, 2, 3, 4, 6, 7},
		},
		{
			[]int{4, 2, 7, 1, 9, 3},
			[]int{1, 2, 3, 4, 7, 9},
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			[]int{1, 2, 3, 4, 6, 7, 9},
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 8},
			[]int{1, 2, 3, 4, 6, 7, 8, 9},
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 10},
			[]int{1, 2, 3, 4, 6, 7, 9, 10},
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 10, 8, 11},
			[]int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11},
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 10, 8, 11, 12},
			[]int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 10, 8, 11, 12, 13, 14},
			[]int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
		{
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		list := inorderTraversal(root)
		fmt.Println(common.IntEqual(test.result, list))
	}
}
