package treeTag

import (
	"fmt"
)

/**
110. 平衡二叉树
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left -right >= 2 || right - left >= 2 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalancedEx(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left -right >= 2 || right - left >= 2 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func TestisBalanced() {
	tests := []struct {
		nums   []int
		result bool
	}{
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, -1, -1, 5, 5},
			true,
		},
		{
			[]int{},
			true,
		},
		{
			[]int{10},
			true,
		},
		{
			[]int{10, 5},
			true,
		},
		{
			[]int{10, 5, 3},
			true,
		},
		{
			[]int{10, 5, 3, 12},
			true,
		},
		{
			[]int{10, 5, 3, 1, 12},
			true,
		},
		{
			[]int{10, 15, 20},
			true,
		},
		{
			[]int{10, 15, 20, 5},
			true,
		},
		{
			[]int{10, 15, 20, 25, 5},
			true,
		},
		{
			[]int{10, 15, 20, 25, 5, 3},
			true,
		},
		{
			[]int{10, 15, 20, 25, 5, 3, 1},
			true,
		},
		{
			[]int{10, 15, 12, 20, 25, 5, 3, 1},
			true,
		},
		{
			[]int{1,2,-1},
			true,
		},
		{
			[]int{1,2,-1,3},
			false,
		},
		{
			[]int{1,2,-1,3,3},
			false,
		},
		{
			[]int{1,2,2,3,3,-1,-1},
			true,
		},
		{
			[]int{1,2,2,3,3,-1,-1,4},
			false,
		},
		{
			[]int{1,2,2,3,3,-1,-1,4,4},
			false,
		},
	}
	for _, test := range tests {
		t1 := ArrayToTree(test.nums)
		result := isBalanced(t1)
		fmt.Println(result == test.result)
	}

}
