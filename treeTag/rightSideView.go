package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
199. 二叉树的右视图
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：从右往左望，看到的是i*2 + 2，如果这个元素不再，就在这一层从右往左找，一定能找到
*/

const empty = -(1<<63 - 1)

func rightSideView(root *TreeNode) []int {
	array := make([]int, 0, 64)
	_rightSideView(root, 0, &array)
	var result []int
	i := 0
	for ; i < len(array); i = i*2 + 2 {
		for index := i; index >= 0; index-- {
			if array[index] != empty {
				result = append(result, array[index])
				break
			}
		}
	}
	if 2*len(array) > i {
		for index := len(array) - 1; index >= 0; index-- {
			if array[index] != empty {
				result = append(result, array[index])
				break
			}
		}
	}
	return result
}

func _rightSideView(root *TreeNode, index int, array *[]int) {
	if root != nil {
		if index >= len(*array) {
			for i := len(*array); i <= index; i++ {
				*array = append(*array, empty)
			}
		}
		(*array)[index] = root.Val
		_rightSideView(root.Left, index*2+1, array)
		_rightSideView(root.Right, index*2+2, array)
	}
}

func TestrightSideView() {
	tests := []struct {
		num    []int
		result []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[]int{1, 3, 7, 15, 16},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			[]int{1, 3, 7, 15},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			[]int{1, 3, 7, 14},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			[]int{1, 3, 7, 13},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			[]int{1, 3, 7, 12},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]int{1, 3, 7, 11},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{1, 3, 7, 10},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 3, 7, 9},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{1, 3, 7, 8},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 3, 7},
		},
		{
			[]int{1, 2, 3, -1, 5, -1, 4},
			[]int{1, 3, 4},
		},
		{
			[]int{1, -1, 3, -1, -1, -1, 4},
			[]int{1, 3, 4},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 3},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{1,},
			[]int{1,},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.num)
		result := rightSideView(root)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
