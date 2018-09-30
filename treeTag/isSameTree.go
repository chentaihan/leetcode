package treeTag

import "fmt"

/**
100. 相同的树
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil || q == nil:
		return q == p
	case p.Val != q.Val:
		return false
	default:
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}

func TestisSameTree() {
	tests := []struct {
		nums1  []int
		nums2  []int
		result bool
	}{
		{
			[]int{},
			[]int{},
			true,
		},
		{
			[]int{10},
			[]int{10},
			true,
		},
		{
			[]int{10, 3},
			[]int{10, 3},
			true,
		},
		{
			[]int{10, 3, 12},
			[]int{10, 3, 12},
			true,
		},
		{
			[]int{4, 1, 2, 0},
			[]int{4, 1, 2},
			false,
		},
	}
	for _, test := range tests {
		t1 := CreateTree(test.nums1)
		t2 := CreateTree(test.nums2)
		result := isSameTree(t1, t2)
		fmt.Println(result == test.result)
	}
}
