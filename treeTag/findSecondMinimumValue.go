package treeTag

import "fmt"

/**
671. 二叉树中第二小的节点
*/

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	queue := Queue{}
	minValue := root.Val
	retVal := -1
	queue.Push(root)
	for !queue.Empty() {
		node := queue.Pop()
		if node.Val != minValue {
			if retVal == -1 {
				retVal = node.Val
			} else if node.Val < retVal {
				retVal = node.Val
			}
		}
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
	}
	return retVal
}

func findSecondMinimumValueEx(root *TreeNode) int {
	first, second := -1, -1
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root != nil {
			if first == -1 {
				first = root.Val
			} else if second == -1 {
				if first != root.Val {
					if first < root.Val {
						second = root.Val
					} else {
						second, first = first, root.Val
					}
				}
			} else {
				if root.Val < second && root.Val != first {
					second = root.Val
					if second < first {
						first, second = second, first
					}
				}
			}
			f(root.Left)
			f(root.Right)
		}
	}
	f(root)
	return second
}

func TestfindSecondMinimumValue() {
	tests := []struct {
		nums   []int
		result int
	}{
		//{
		//	[]int{1},
		//	-1,
		//},
		{
			[]int{5, 8, 5},
			8,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 2},
			2,
		},
		{
			[]int{1, 2, 3},
			2,
		},
		{
			[]int{2, 2, 2},
			-1,
		},
		{
			[]int{2, 2, 3},
			3,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 3},
			3,
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2},
			-1,
		},
		{
			[]int{2, 2, 2, 2, 4, 2, 3},
			3,
		},
		{
			[]int{2, 2, 5, -1, -1, 5, 7},
			5,
		},
		{
			[]int{2, 2, 5, -1, -1, 6, 7},
			5,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := findSecondMinimumValueEx(root)
		fmt.Println(test.result == result)
	}
}
