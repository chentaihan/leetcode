package depthFirstTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/treeTag"
	"math"
)

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
//

/**
最小节点
*/
func MinNode(root *treeTag.TreeNode) *treeTag.TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

/**
最大节点
*/
func MaxNode(root *treeTag.TreeNode) *treeTag.TreeNode {
	if root == nil {
		return root
	}
	if root.Right == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func isValidBST(root *treeTag.TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil  {
		if root.Val <= root.Left.Val {
			return false
		}
		node := MaxNode(root.Left)
		if node != nil && node.Val >= root.Val {
			return false
		}
	}

	if root.Right != nil  {
		if root.Val >= root.Right.Val  {
			return false
		}
		node := MinNode(root.Right)
		if node != nil && node.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}


func isValidBSTEx(root *treeTag.TreeNode) bool {
	return isValidBSTUtil(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTUtil(root *treeTag.TreeNode, min,max int) bool{
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValidBSTUtil(root.Left, min, root.Val) && isValidBSTUtil(root.Right, root.Val, max)
}


func TestisValidBST() {
	tests := []struct {
		nums   []int
		result bool
	}{
		{
			[]int{4},
			true,
		},
		{
			[]int{4, 1},
			true,
		},
		{
			[]int{4, 1, 2},
			false,
		},
		{
			[]int{4, 1, 5},
			true,
		},
		{
			[]int{4, 1, 5, -1, 2, -1, 7},
			true,
		},
		{
			[]int{4, 1, 10, -1, 2, 3, 11},
			false,
		},
		{
			[]int{4, 1, 10, -1, 2, 6, 11},
			true,
		},
	}
	for _, test := range tests {
		root := treeTag.ArrayToTree(test.nums)
		val := isValidBSTEx(root)
		fmt.Println(test.result == val)
	}
}
