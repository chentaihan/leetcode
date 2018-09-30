package treeTag

import "fmt"

/**
563. 二叉树的坡度
给定一个二叉树，计算整个树的坡度。
一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。
整个树的坡度就是其所有节点的坡度之和。
*/

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	totalCount := 0
	_findTilt(root, &totalCount)
	return totalCount
}

func _findTilt(root *TreeNode, totalCount *int) {
	if root != nil {
		leftCount := TreeNodeValue(root.Left)
		rightCount := TreeNodeValue(root.Right)
		if leftCount > rightCount {
			*totalCount += leftCount - rightCount
		} else {
			*totalCount += rightCount - leftCount
		}
		_findTilt(root.Left, totalCount)
		_findTilt(root.Right, totalCount)
	}
}

func TreeNodeValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + TreeNodeValue(root.Left) + TreeNodeValue(root.Right)
}

func TestfindTilt() {
	tests := []struct {
		nums       []int
		totalCount int
	}{
		{
			[]int{1},
			0,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 2},
			0,
		},
		{
			[]int{1, 2, 2, 3},
			6,
		},
		{
			[]int{1, 2, 2, 3, 3},
			6,
		},
		{
			[]int{1, 2, 2, 3, 3, 3},
			6,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3},
			0,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4},
			12,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4},
			16,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, -1, -1, -1, 4},
			16,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4},
			0,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5},
			20,
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5},
			30,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		totalCount := findTilt(root)
		fmt.Println(test.totalCount == totalCount)
	}
}
