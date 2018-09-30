package treeTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
226. 翻转二叉树
输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
 */

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

func TestinvertTree() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			[]int{4, 7, 9, 6, 2, 3, 1},
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		invertTree(root)
		list := []int{}
		DLRList(root, &list)
		fmt.Println(common.IntEqual(test.result, list))
	}
}
