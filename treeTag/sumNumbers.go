package treeTag

import "fmt"

//给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
//
// 例如，从根到叶子节点路径 1->2->3 代表数字 123。
//
// 计算从根到叶子节点生成的所有数字之和。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例 1:
//
// 输入: [1,2,3]
//    1
//   / \
//  2   3
//输出: 25
//解释:
//从根到叶子节点路径 1->2 代表数字 12.
//从根到叶子节点路径 1->3 代表数字 13.
//因此，数字总和 = 12 + 13 = 25.
//
// 示例 2:
//
// 输入: [4,9,0,5,1]
//    4
//   / \
//  9   0
// / \
//5   1
//输出: 1026
//解释:
//从根到叶子节点路径 4->9->5 代表数字 495.
//从根到叶子节点路径 4->9->1 代表数字 491.
//从根到叶子节点路径 4->0 代表数字 40.
//因此，数字总和 = 495 + 491 + 40 = 1026.
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	sum := 0
	sumNumbersUtil(root, 0, &sum)
	return sum
}

func sumNumbersUtil(root *TreeNode, value int, sum *int) {
	if root != nil {
		value = value*10 + root.Val
		if root.Left == nil && root.Right == nil {
			*sum += value
		}
		sumNumbersUtil(root.Left, value, sum)
		sumNumbersUtil(root.Right, value, sum)
	}
}

func sumNumbersEx(root *TreeNode) int {
	sum := 0
	var f func(root *TreeNode, value int)
	f = func(root *TreeNode, value int) {
		if root == nil {
			return
		}
		value = value*10 + root.Val
		if root.Left == nil && root.Right == nil {
			sum += value
		}
		f(root.Left, value)
		f(root.Right, value)
	}
	f(root, 0)
	return sum
}

func TestsumNumbers() {
	tests := []struct {
		list []int
		sum  int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2},
			12,
		},
		{
			[]int{1, 2, 3},
			25,
		},
		{
			[]int{1, 2, 3, 4},
			137,
		},
		{
			[]int{1, 2, 3, 4, 5},
			262,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			385,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			522,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			1646,
		},
	}
	for _, test := range tests {
		head := ArrayToTree(test.list)
		sum := sumNumbersEx(head)

		fmt.Println(sum == test.sum)
	}
}
