package treeTag

import "fmt"

/**
662. 二叉树最大宽度
 */

type ValueItem struct {
	minValue int
	maxValue int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	array := []*ValueItem{}
	_widthOfBinaryTree(root, 0, 0, &array)
	result := 0
	for _, item := range array {
		if item.maxValue-item.minValue > result {
			result = item.maxValue - item.minValue
		}
	}
	return result + 1
}

func _widthOfBinaryTree(root *TreeNode, index, level int, array *[]*ValueItem) {
	if root != nil {
		for len(*array) <= level {
			*array = append(*array, &ValueItem{})
		}
		curItem := (*array)[level]
		if curItem.minValue == curItem.maxValue && curItem.minValue == 0 {
			curItem.minValue, curItem.maxValue = index, index
		} else {
			if curItem.minValue > index {
				curItem.minValue = index
			}
			if curItem.maxValue < index {
				curItem.maxValue = index
			}
		}
		level++
		_widthOfBinaryTree(root.Left, index*2+1, level, array)
		_widthOfBinaryTree(root.Right, index*2+2, level, array)
	}
}

func TestwidthOfBinaryTree() {
	tests := []struct {
		nums  []int
		width int
	}{
		{
			[]int{4},
			1,
		},
		{
			[]int{4, 2},
			1,
		},
		{
			[]int{4, 7, 9},
			1,
		},
		{
			[]int{4, 1, 7, 9},
			2,
		},
		{
			[]int{4, 3, 7, 1},
			2,
		},
		{
			[]int{4, 2, 7, 1, 3},
			2,
		},
		{
			[]int{4, 2, 7, 1, 3, 9},
			4,
		},
		{
			[]int{4, 2, 7, 1, 3, 6},
			3,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			4,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 10},
			4,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9, 10, 12},
			4,
		},
		{
			[]int{20, 10, 30, 5, 15, 25, 40, 2, 6, 8, 12, 22, 28, 35, 50},
			8,
		},
		{
			[]int{20, 10, 30, 5, 15, 25, 40, 2, 6, 8, 12, 22, 28, 35, 50, 60, 1},
			16,
		},
		{
			[]int{20, 10, 30, 5, 15, 25, 40, 2, 6, 8, 12, 22, 28, 35, 50, 45, 1},
			15,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		width := widthOfBinaryTree(root)
		fmt.Println(test.width == width)
	}
}
