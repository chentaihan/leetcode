package treeTag

import "fmt"

/**
437. 路径总和 III
*/

func pathSum(root *TreeNode, sum int) int {
	if root != nil {
		return _pathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
	}
	return 0
}

func _pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Val == sum {
		return 1 + _pathSum(root.Left, sum-root.Val) + _pathSum(root.Right, sum-root.Val)
	}
	return _pathSum(root.Left, sum-root.Val) + _pathSum(root.Right, sum-root.Val)
}

func TestpathSum() {
	tests := []struct {
		nums  []int
		sum   int
		count int
	}{
		{
			[]int{},
			0,
			0,
		},
		{
			[]int{2},
			2,
			1,
		},
		{
			[]int{2, 2},
			2,
			2,
		},
		{
			[]int{2, 2, 2},
			2,
			3,
		},
		{
			[]int{2, 2, 2},
			4,
			2,
		},
		{
			[]int{2, 2, 2, 2},
			4,
			3,
		},
		{
			[]int{2, 2, 2, 2, 2},
			4,
			4,
		},
		{
			[]int{2, 2, 2, 2, 2, 2},
			4,
			5,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2},
			4,
			6,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2},
			4,
			7,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2},
			4,
			8,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			6,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			7,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			8,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			9,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			10,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			11,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			6,
			12,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			8,
			8,
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		count := pathSum(root, test.sum)
		fmt.Println(count == test.count)
	}
}
