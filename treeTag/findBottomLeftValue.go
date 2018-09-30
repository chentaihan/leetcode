package treeTag

import "fmt"

/**
513. 找树左下角的值
给定一个二叉树，在树的最后一行找到最左边的值。
 */

func findBottomLeftValue(root *TreeNode) int {
	val := 0
	maxLevel := -1
	_findBottomLeftValue(root, 0, &maxLevel, &val)
	return val
}

func _findBottomLeftValue(root *TreeNode, level int, maxLevel *int, val *int) {
	if root != nil {
		if level > *maxLevel {
			*maxLevel = level
			*val = root.Val
		}
		level++
		_findBottomLeftValue(root.Left, level, maxLevel, val)
		_findBottomLeftValue(root.Right, level, maxLevel, val)
	}
}

func TestfindBottomLeftValue() {
	tests := []struct {
		nums []int
		val  int
	}{
		{
			[]int{4},
			4,
		},
		{
			[]int{4, 1},
			1,
		},
		{
			[]int{4, 1, 2},
			1,
		},
		{
			[]int{4, 1, 2, 3},
			3,
		},
		{
			[]int{4, 1, 2, 3, 5},
			3,
		},
		{
			[]int{4, 1, 2, 3, 5, 6},
			3,
		},
		{
			[]int{4, 1, 2, 3, 5, 6, 7},
			3,
		},
		{
			[]int{4, 1, 2, 3, 5, 6, 7, -1, -1, 8},
			8,
		},
		{
			[]int{4, 1, 2, 3, 5, 6, 7, 8, -1, 9, -1, 10},
			8,
		},
		{
			[]int{4, 1, 2, 3, 5, 6, 7, -1, -1, 8, -1, -1, -1, -1, -1, -1, -1, -1, -1, 9},
			9,
		},
		{
			[]int{1, 2, 3, 4, -1, 5, 6, -1, -1, -1, -1, 7},
			7,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		val := findBottomLeftValue(root)
		fmt.Println(test.val == val)
	}
}
