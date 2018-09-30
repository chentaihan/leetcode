package treeTag

import "fmt"

/**
230. 二叉搜索树中第K小的元素
 */

func kthSmallest(root *TreeNode, k int) int {
	result := 0
	_kthSmallest(root, &k, &result)
	return result
}

func _kthSmallest(root *TreeNode, k *int, result *int) {
	if root != nil {
		_kthSmallest(root.Left, k, result)
		*k--
		if *k == 0 {
			*result = root.Val
		} else {
			_kthSmallest(root.Right, k, result)
		}
	}
}

func kthSmallestEx(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack := NewStackTree()
	stack.Push(root)
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
			k--
			if k == 0 {
				return root.Val
			}
			//右子树
			if root.Right != nil {
				stack.Push(root.Right)
				root = root.Right
				break
			}
		}
	}
	return 0
}

func TestkthSmallest() {
	tests := []struct {
		nums []int
		k    int
		val  int
	}{
		{
			[]int{4},
			1,
			4,
		},
		{
			[]int{4, 2},
			2,
			4,
		},
		{
			[]int{4, 7, 9},
			3,
			9,
		},
		{
			[]int{4, 1, 7, 9},
			3,
			7,
		},
		{
			[]int{4, 3, 7, 1},
			2,
			3,
		},
		{
			[]int{4, 2, 7, 1, 3},
			1,
			1,
		},
		{
			[]int{4, 2, 7, 1, 3, 9},
			2,
			2,
		},
		{
			[]int{4, 2, 7, 1, 3, 6},
			3,
			3,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			4,
			4,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			5,
			6,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			6,
			7,
		},
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			7,
			9,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		val := kthSmallestEx(root, test.k)
		fmt.Println(test.val == val)
	}
}
