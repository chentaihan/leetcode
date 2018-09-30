package treeTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
654. 最大二叉树
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。

Example 1:

输入: [3,2,1,6,0,5]
输入: 返回下面这棵树的根节点：

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := findMaxValue(nums, 0, len(nums))
	root := &TreeNode{
		Val: nums[index],
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func findMaxValue(nums []int, start, end int) int {
	index := start
	for i := start + 1; i < end; i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}

func TestconstructMaximumBinaryTree() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{2, 1},
			[]int{2, 1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{1, 3, 2},
			[]int{3, 1, 2},
		},
		{
			[]int{2, 3, 1},
			[]int{3, 2, 1},
		},
		{
			[]int{2, 4, 1, 3},
			[]int{4, 2, 3, 1},
		},
		{
			[]int{4, 2, 1, 3},
			[]int{4, 3, 2, 1},
		},
		{
			[]int{3, 2, 1, 6, 0, 5},
			[]int{6, 3, 5, 2, 0, 1},
		},
		{
			[]int{3, 2, 7, 1, 6, 0, 5},
			[]int{7, 3, 6, 2, 1, 5, 0},
		},
	}
	for _, test := range tests {
		root := constructMaximumBinaryTree(test.nums)
		result := TreeFloor(root)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
