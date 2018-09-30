package treeTag

/**
111. 二叉树的最小深度
 */

import (
	"fmt"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := QueueEx{}
	queue.Push(root, 1)
	for !queue.Empty() {
		node, level := queue.Pop()
		if node.Left == nil && node.Right == nil {
			return level
		}
		if node.Left != nil {
			queue.Push(node.Left, level+1)
		}
		if node.Right != nil {
			queue.Push(node.Right, level+1)
		}
	}
	return 0
}

func minDepthEx(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepthEx(root.Right)
	case root.Right == nil:
		return 1 + minDepthEx(root.Left)
	default:
		return 1 + min(minDepthEx(root.Left), minDepthEx(root.Right))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinDepth() {
	tests := []struct {
		nums  []int
		depth int
	}{
		{
			[]int{4},
			1,
		},
		{
			[]int{4, 2},
			2,
		},
		{
			[]int{4, 7, 9},
			3,
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
			3,
		},
		{
			[]int{4, 2, 7, 1, 3, 6},
			3,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			3,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		depth := minDepth(root)
		fmt.Println(test.depth == depth)
	}
}
