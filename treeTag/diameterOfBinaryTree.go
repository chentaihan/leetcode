package treeTag

import "fmt"

/**
543. 二叉树的直径
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxVal := 0
	diameterTree(root, &maxVal)
	return maxVal
}

func diameterTree(root *TreeNode, maxVal *int) {
	if root != nil {
		val := maxDepth(root.Left) + maxDepth(root.Right)
		if val > *maxVal {
			*maxVal = val
		}
		diameterTree(root.Left, maxVal)
		diameterTree(root.Right, maxVal)
	}
}

func TestdiameterOfBinaryTree() {
	tests := []struct {
		nums  []int
		depth int
	}{
		{
			[]int{4},
			0,
		},
		{
			[]int{4, 2},
			1,
		},
		{
			[]int{4, 7, 9},
			2,
		},
		{
			[]int{4, 1, 7, 9},
			3,
		},
		{
			[]int{4, 3, 7, 1},
			3,
		},
		{
			[]int{4, 2, 7, 1, 3},
			3,
		},
		{
			[]int{4, 2, 7, 1, 3, 9},
			4,
		},
		{
			[]int{4, 2, 7, 1, 3, 6},
			4,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			4,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		depth := diameterOfBinaryTree(root)
		fmt.Println(test.depth == depth)
	}
}
