
package treeTag

import "fmt"

//给定一个二叉树，原地将它展开为链表。
//
// 例如，给定二叉树
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6
//
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = nil
	if left != nil {
		root.Right = left
		flatten(left)
	}
	node := root
	for node.Right != nil {
		node = node.Right
	}
	node.Right = right
	flatten(right)
}

func Testflatten() {
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
			[]int{1, 2, 5, 3, 4, -1, 6},
		},
		{
			[]int{1, -1, 2, -1, -1, -1, 3},
		},
		{
			[]int{1, 2, 3, -1, -1, 4, 5},
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
		head := ArrayToTree(test.list)
		flatten(head)
		for head != nil {
			fmt.Print(head.Val, " ")
			head = head.Right
		}
		fmt.Println()
	}
}
