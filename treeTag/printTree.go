package treeTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
	"strconv"
)

/**
655. 输出二叉树
分析：变体的二分查找
*/

func printTree(root *TreeNode) [][]string {
	depth := TreeDepth(root)
	result := make([][]string, depth)
	colCount := 1<<uint(depth) - 1
	for i := 0; i < depth; i++ {
		result[i] = make([]string, colCount)
	}
	var _printTree func(root *TreeNode, curDepth uint, start, end int)
	_printTree = func(root *TreeNode, curDepth uint, start, end int) {
		if root != nil {
			curIndex := (start + end) / 2
			result[curDepth][curIndex] = strconv.Itoa(root.Val)
			curDepth++
			_printTree(root.Left, curDepth, start, curIndex-1)
			_printTree(root.Right, curDepth, curIndex+1, end)
		}
	}
	_printTree(root, 0, 0, colCount-1)
	return result
}

func TestprintTree() {
	tests := []struct {
		nums   []int
		result [][]string
	}{
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			[][]string{{"", "", "", "5", "", "", ""}, {"", "3", "", "", "", "6", ""}, {"2", "", "4", "", "", "", "7"}},
		},
		{
			[]int{1},
			[][]string{{"1"}},
		},
		{
			[]int{1, 2},
			[][]string{{"", "1", ""}, {"2", "", ""}},
		},
		{
			[]int{1, 2, 2},
			[][]string{{"", "1", ""}, {"2", "", "2"}},
		},
		{
			[]int{1, 2, 2, 3},
			[][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "2", ""}, {"3", "", "", "", "", "", ""}},
		},

		{
			[]int{1, 2, 2, 3, 3},
			[][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "2", ""}, {"3", "", "3", "", "", "", ""}},
		},
		{
			[]int{1, 2, 3, -1, 4},
			[][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "3", ""}, {"", "", "4", "", "", "", ""}},
		},
		{
			[]int{1, 2, 5, 3, -1, -1, -1, 4},
			[][]string{{"", "", "", "", "", "", "", "1", "", "", "", "", "", "", ""}, {"", "", "", "2", "", "", "", "", "", "", "", "5", "", "", ""}, {"", "3", "", "", "", "", "", "", "", "", "", "", "", "", ""}, {"4", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := printTree(root)
		fmt.Println(common.StringArrayEqualEx(test.result, result))
	}
}
