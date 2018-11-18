package treeTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树

分析：按照前序遍历，相同的值会连在一起，统计连续的数量，肯定比将每个元素直接map性能好
*/

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := NewStackTree()
	stack.Push(root)
	list := []int{}
	count, maxCount, curVal := 0, 0, root.Val
	for !stack.Empty() {
		//左子树，一步到叶子节点
		for root.Left != nil {
			stack.Push(root.Left)
			root = root.Left
		}
		//从叶子节点往根节点迭代
		for !stack.Empty() {
			//根节点
			root = stack.Pop()
			if root.Val == curVal {
				count++
			} else {
				count = 1
				curVal = root.Val
			}

			if count > maxCount {
				list = append(list[:0], root.Val)
				maxCount = count
			} else if count == maxCount && (len(list) == 0 || list[len(list)-1] != root.Val ) {
				list = append(list, root.Val)
			}

			//右子树
			if root.Right != nil {
				stack.Push(root.Right)
				root = root.Right
				break
			}
		}
	}
	return list
}
 
func TestfindMode() {
	tests := []struct {
		nums []int
		ret  []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			[]int{2, -1, 3},
			[]int{2, 3},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{2, 2, 3},
			[]int{2},
		},
		{
			[]int{3, 2, 3},
			[]int{3},
		},
		{
			[]int{3, 3, 3},
			[]int{3},
		},
		{
			[]int{1, -1, 2, -1, -1, 2},
			[]int{2},
		},
		{
			[]int{1, 2, 3, -1, -1, 3, 3},
			[]int{3},
		},
		{
			[]int{1, 1, 3, -1, -1, 3},
			[]int{1, 3},
		},
		{
			[]int{1, 1, 3, -1, -1, 3, 3},
			[]int{3},
		},
		{
			[]int{1, 1, 1, -1, -1, 3, 3},
			[]int{1},
		},
		{
			[]int{1, 1, 1, -1, -1, 3, 3, -1, -1, -1, -1, 3},
			[]int{1, 3},
		},
		{
			[]int{1, 1, 3, 1, -1, 3, 3, -1, -1, -1, -1, 3, 3},
			[]int{3},
		},
		{
			[]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 2, 6},
			[]int{2, 6},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		ret := findMode(root)
		fmt.Println(common.IntEqualSort(ret, test.ret))
	}

}
