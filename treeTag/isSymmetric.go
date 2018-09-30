package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
101. 对称二叉树
*/
func isSymmetric(root *TreeNode) bool {
	if root != nil {
		return _isSymmetric(root.Left, root.Right)
	}
	return true
}

func _isSymmetric(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil || q == nil:
		return q == p
	case p.Val != q.Val:
		return false
	default:
		return _isSymmetric(p.Left, q.Right) && _isSymmetric(p.Right, q.Left)
	}
}

func TestisSymmetric() {
	tests := []struct {
		nums   []int
		floor  []int
		result bool
	}{
		{
			[]int{},
			[]int{},
			true,
		},
		{
			[]int{2},
			[]int{2},
			true,
		},
		{
			[]int{2, 3, 3},
			[]int{2, 3, 3},
			true,
		},
		{
			[]int{2, 3, 3, 4, 5, -1, 4},
			[]int{2, 3, 3, 4, 5, 4},
			false,
		},
		{
			[]int{2, 3, 3, 4, 5, 5, 4},
			[]int{2, 3, 3, 4, 5, 5, 4},
			true,
		},
		{
			[]int{2, 3, 3, 4, 5, 6, 4},
			[]int{2, 3, 3, 4, 5, 6, 4},
			false,
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		floor := TreeFloor(root)
		result := isSymmetric(root)
		fmt.Println(result == test.result)
		fmt.Println(common.IntEqual(floor, test.floor))
	}
}
