package treeTag

import "fmt"

/**
637. 二叉树的层平均值
 */

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	if root == nil {
		return ret
	}
	queue := QueueEx{}
	curLevel := 0
	sum := 0
	queue.Push(root, 0)
	count := 0
	for !queue.Empty() {
		node, level := queue.Pop()
		if curLevel != level {
			curLevel = level
			ret = append(ret, float64(sum)/float64(count))
			sum = 0
			count = 0
		}
		sum += node.Val
		count++
		if node.Left != nil {
			queue.Push(node.Left, level+1)
		}
		if node.Right != nil {
			queue.Push(node.Right, level+1)
		}
	}
	ret = append(ret, float64(sum)/float64(count))
	return ret
}

func TestaverageOfLevels() {
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
			[]int{1, 3},
		},
		{
			[]int{2, 2, 2},
			[]int{2, 2},
		},
		{
			[]int{2, 2, 4, 4, 6},
			[]int{2, 3, 5},
		},
		{
			[]int{2, 2, 4, -1, -1, 4, 6},
			[]int{2, 3, 5},
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 6},
			[]int{2, 2, 3},
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2},
			[]int{2, 2, 2},
		},
		{
			[]int{2, 2, 2, 2, 4, 2, 4, 4},
			[]int{2, 2, 3, 4},
		},
		{
			[]int{2, 2, 6, -1, -1, 5, 7},
			[]int{2, 4, 6},
		},
		{
			[]int{2, 2, 6, -1, -1, 7, 7, 2, 4, 6, 8},
			[]int{2, 4, 7, 5},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := averageOfLevels(root)
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
