package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
113. 路径总和 II
*/

func pathSum2(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	array := []int{}
	_pathSum2(root, sum, array, &result)
	return result
}

func _pathSum2(root *TreeNode, sum int, array []int, result *[][]int) {
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			*result = append(*result, append(array, root.Val))
		}
	}
	if root.Left != nil || root.Right != nil {
		array = append(array, root.Val)
		sum -= root.Val
		if root.Left != nil {
			_pathSum2(root.Left, sum, array, result)
			if root.Right != nil {
				rightArray := make([]int, len(array))
				copy(rightArray, array)
				_pathSum2(root.Right, sum, rightArray, result)
			}
		} else {
			_pathSum2(root.Right, sum, array, result)
		}
	}
}

func TestpathSum2() {
	tests := []struct {
		nums   []int
		sum    int
		result [][]int
	}{
		{
			[]int{1},
			1,
			[][]int{{1}},
		},
		{
			[]int{1, 2, 2},
			3,
			[][]int{{1, 2}, {1, 2}},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3},
			6,
			[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4},
			10,
			[][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}},
		},
		{
			[]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1},
			22,
			[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := pathSum2(root, test.sum)
		fmt.Println(common.IntEqualEx(result, test.result))
	}
}
