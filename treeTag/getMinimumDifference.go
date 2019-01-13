package treeTag

import (
	"fmt"
)

/**
530. 二叉搜索树的最小绝对差
给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

示例 :

输入:
   1
    \
     3
    /
   2

输出:
1

解释:
最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
注意: 树中至少有2个节点。
*/

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func getMinimumDifference(root *TreeNode) int {
	minValue := 1<<63 - 1
	_getMinimumDifference(root, &minValue)
	return minValue
}

func _getMinimumDifference(root *TreeNode, minValue *int) {
	if root != nil {
		tmpNode := root.Left
		//左子树最大值
		if tmpNode != nil {
			for tmpNode.Right != nil {
				tmpNode = tmpNode.Right
			}
			if root.Val- tmpNode.Val < *minValue {
				*minValue = root.Val- tmpNode.Val
			}
		}

		tmpNode = root.Right
		//右子树最小值
		if tmpNode != nil {
			for tmpNode.Left != nil {
				tmpNode = tmpNode.Left
			}
			if tmpNode.Val-root.Val < *minValue {
				*minValue = tmpNode.Val - root.Val
			}
		}

		_getMinimumDifference(root.Left, minValue)
		_getMinimumDifference(root.Right, minValue)
	}
}

func TestgetMinimumDifference() {
	tests := []struct {
		value  []int
		result int
	}{
		{
			[]int{5, 1},
			4,
		},
		{
			[]int{1, -1, 5},
			4,
		},
		{
			[]int{5, 1, 7},
			2,
		},
		{
			[]int{5, 1, 11},
			4,
		},
		{
			[]int{5, 1, 7, -1, 2},
			1,
		},
		{
			[]int{5, 3, 11, 1, -1, 9},
			2,
		},
		{
			[]int{5, 1, 7, 1, 2},
			0,
		},
		{
			[]int{5, 1, 11, -1, 4, 9},
			1,
		},
		{
			[]int{5, 1, 10, -1, 3, 8, 12},
			2,
		},
		{
			[]int{236, 104, 701, -1, 227, -1, 911},
			9,
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.value)
		result := getMinimumDifference(root)
		fmt.Println(result == test.result)
	}
}
