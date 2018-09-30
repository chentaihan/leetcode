package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
108. 将有序数组转换为二叉搜索树
二分搜索数组，构造二叉树
 */

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	binaryArray(nums, 0, len(nums)-1, root)
	return root
}

func binaryArray(arr []int, start, end int, root *TreeNode) {
	middle := (start + end ) / 2
	root.Val = arr[middle]
	if start <= middle-1 {
		root.Left = &TreeNode{}
		binaryArray(arr, start, middle-1, root.Left)
	}
	if middle+1 <= end {
		root.Right = &TreeNode{}
		binaryArray(arr, middle+1, end, root.Right)
	}
}

/**
膜拜版本
 */
func sortedArrayToBSTEx(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) >> 1
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBSTEx(nums[:m]),
		Right: sortedArrayToBSTEx(nums[m+1:]),
	}
}

//构造测试结果需要的数据
func binaryArrayTest(arr []int, start, end int, root *TreeNode) {
	middle := (start + end ) / 2
	fmt.Print(arr[middle], ", ")
	if start <= middle-1 {
		binaryArrayTest(arr, start, middle-1, root)
	}
	if middle+1 <= end {
		binaryArrayTest(arr, middle+1, end, root)
	}
}

func TestsortedArrayToBST() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 1, 2, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{3, 1, 2, 5, 4, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 1, 3, 6, 5, 7},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{4, 2, 1, 3, 6, 5, 7, 8},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{4, 1, 0, 2, 3, 7, 5, 6, 8, 9},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{5, 2, 0, 1, 3, 4, 8, 6, 7, 9, 10},
		},
	}

	for _, test := range tests {
		root := sortedArrayToBST(test.nums)
		result := []int{}
		DLRList(root, &result)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
