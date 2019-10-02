package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,2,3]
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var list []int
	preorderTraversalUtil(root, &list)
	return list
}

func preorderTraversalUtil(root *TreeNode, list *[]int) {
	if root != nil {
		*list = append(*list, root.Val)
		preorderTraversalUtil(root.Left, list)
		preorderTraversalUtil(root.Right, list)
	}
}

func TestpreorderTraversal() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, -1, 3},
			[]int{1, 3},
		},
		{
			[]int{1,},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 4, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 4, 5, 3},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 4, 5, 3, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 4, 5, 3, 6, 7},
		},
	}

	for _,test := range tests {
		root := ArrayToTree(test.nums)
		result := preorderTraversal(root)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
