package treeTag

/**
112. 路径总和
 */

import (
	"fmt"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

/**
膜拜版本，switch case还是让我眼前一亮啊
 */
func hasPathSumEx(root *TreeNode, sum int) bool {
	switch {
	case root == nil:
		return false
	case root.Left == nil && root.Right == nil:
		return root.Val == sum
	default:
		return hasPathSumEx(root.Left, sum-root.Val) || hasPathSumEx(root.Right, sum-root.Val)
	}
}

func TesthasPathSum() {
	tests := []struct {
		nums   []int
		sum    int
		result bool
	}{
		{
			[]int{},
			0,
			false,
		},
		{
			[]int{1},
			0,
			false,
		},
		{
			[]int{2},
			2,
			true,
		},
		{
			[]int{1, 2},
			1,
			false,
		},
		{
			[]int{1, 2},
			3,
			true,
		},
		{
			[]int{2, 1},
			3,
			true,
		},
		{
			[]int{1, 2},
			4,
			false,
		},
		{
			[]int{4, 2, 7},
			6,
			true,
		},
		{
			[]int{4, 2, 7},
			11,
			true,
		},
		{
			[]int{4, 2, 7},
			12,
			false,
		},
		{
			[]int{2, 4, 7},
			13,
			true,
		},
		{
			[]int{2, 4, 7},
			15,
			false,
		},
		{
			[]int{4, 2, 7, 1},
			7,
			true,
		},
		{
			[]int{4, 2, 7, 1},
			11,
			true,
		},
		{
			[]int{4, 2, 7, 1},
			71,
			false,
		},
		{
			[]int{4, 2, 7, 1},
			71,
			false,
		},
		{
			[]int{4, 2, 7, 1, 5},
			7,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5},
			16,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5},
			17,
			false,
		},
		{
			[]int{4, 2, 7, 1, 5, 3},
			7,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3},
			9,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3},
			7,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3, 8},
			19,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3, 8, 10},
			29,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3, 8, 10, 12},
			41,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3, 8, 10, 12, 19},
			60,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3, 8, 10, 12, 19, 9},
			38,
			true,
		},
		{
			[]int{4, 2, 7, 1, 5, 3, 8, 10, 12, 19, 9, 11},
			52,
			true,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		result := hasPathSum(root, test.sum)
		fmt.Println(test.result == result)
		result = hasPathSumEx(root, test.sum)
		fmt.Println(test.result == result)
	}
}
