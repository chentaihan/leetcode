package treeTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
538. 把二叉搜索树转换为累加树
*/

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	_convertBST(root, &sum)
	return root
}

func _convertBST(root *TreeNode, sum *int) {
	if root != nil {
		_convertBST(root.Right, sum)
		root.Val += *sum
		*sum = root.Val
		_convertBST(root.Left, sum)
	}
}

/**
非递归实现，把二叉搜索树转换为累加树
非递归也不一定比递归快啊
 */
func convertBSTEx(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	curRoot := root
	stack := NewStackTree()
	stack.Push(curRoot)
	sum := 0
	for !stack.Empty() {
		//右子树，一步到叶子节点
		for curRoot.Right != nil {
			stack.Push(curRoot.Right)
			curRoot = curRoot.Right
		}
		//从叶子节点往根节点迭代
		for !stack.Empty() {
			//根节点
			curRoot = stack.Pop()
			curRoot.Val += sum
			sum = curRoot.Val
			//左子树
			if curRoot.Left != nil {
				stack.Push(curRoot.Left)
				curRoot = curRoot.Left
				break
			}
		}
	}
	return root
}



func TestconvertBST() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{3, 2},
		},
		{
			[]int{2, 1},
			[]int{2, 3},
		},
		{
			[]int{2, 1, 3},
			[]int{5, 6, 3},
		},
		{
			[]int{5, 3, 6, 2, 4, 7},
			[]int{18, 25, 13, 27, 22, 7},
		},
		{
			[]int{5, 3, 7, 2, 4, 6, 8},
			[]int{26, 33, 15, 35, 30, 21, 8},
		},
		{
			[]int{5, 3, 7, 2, 4, 6, 8, 1, 9},
			[]int{35, 42, 24, 44, 39, 30, 17, 45, 9},
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		root = convertBSTEx(root)
		result := TreeFloor(root)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
