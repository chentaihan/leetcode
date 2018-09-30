package treeTag

import "fmt"

/**
515. 在每个树行中找最大值
 */

func largestValues(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	queue := QueueEx{}
	curLevel := 0
	maxValue := -1 << 63
	queue.Push(root, 0)
	for !queue.Empty() {
		node, level := queue.Pop()
		if curLevel != level {
			curLevel = level
			ret = append(ret, maxValue)
			maxValue = -1 << 63
		}
		if maxValue < node.Val {
			maxValue = node.Val
		}
		if node.Left != nil {
			queue.Push(node.Left, level+1)
		}
		if node.Right != nil {
			queue.Push(node.Right, level+1)
		}
	}
	ret = append(ret, maxValue)
	return ret
}

func TestlargestValues() {
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
			[]int{1, 2},
		},
		{
			[]int{1, 2, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 4},
			[]int{1, 4},
		},
		{
			[]int{2, 2, 2},
			[]int{2, 2},
		},
		{
			[]int{2, 2, 4, 4, 6},
			[]int{2, 4, 6},
		},
		{
			[]int{2, 2, 4, -1, -1, 4, 6},
			[]int{2, 4, 6},
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 6},
			[]int{2, 2, 6},
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2},
			[]int{2, 2, 2},
		},
		{
			[]int{2, 2, 2, 2, 4, 2, 4, 4},
			[]int{2, 2, 4, 4},
		},
		{
			[]int{2, 2, 6, -1, -1, 5, 7},
			[]int{2, 6, 7},
		},
		{
			[]int{2, 2, 6, -1, -1, 7, 7, 2, 4, 6, 8},
			[]int{2, 6, 7, 8},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := largestValues(root)
		isOk := true
		for index,val := range result{
			if int(val) != test.result[index]{
				isOk = false
				break
			}
		}
		fmt.Println(isOk)
	}

}
