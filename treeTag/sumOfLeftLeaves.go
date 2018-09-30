package treeTag

import "fmt"

/**
404. 左叶子之和
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 */

func sumOfLeftLeaves(root *TreeNode) int {
	return _sumOfLeftLeaves(root, false)
}

func _sumOfLeftLeaves(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if isLeft && root.Left == nil && root.Right == nil {
		return root.Val
	}
	return _sumOfLeftLeaves(root.Left, true) + _sumOfLeftLeaves(root.Right, false)
}

func TestsumOfLeftLeaves() {
	tests := []struct {
		nums []int
		sum  int
	}{
		{
			[]int{1},
			0,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 2},
			2,
		},
		{
			[]int{1, 2, 2, 3},
			3,
		},
		{
			[]int{1, 2, 2, 3, 3},
			3,
		},
		{
			[]int{1, 2, 2, 3, 3, 3},
			6,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3},
			6,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4},
			7,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4},
			7,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, -1, -1, -1, 4},
			8,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4},
			16,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5},
			17,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5},
			17,
		},
		{
			[]int{3, 9, 20, -1, -1, 15, 7},
			24,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		sum := sumOfLeftLeaves(root)
		fmt.Println(test.sum == sum)
	}
}
