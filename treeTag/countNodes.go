package treeTag

import "fmt"

/**
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func countNodesEx(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return _countNodes(root, 1)
}

func _countNodes(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	right, left := count, count
	if root.Right != nil {
		right = _countNodes(root.Right, 2*count+1)
	}
	if root.Left != nil {
		left = _countNodes(root.Left, 2*count)
	}
	if left > right {
		return left
	}
	return right
}

func TestcountNodes() {
	tests := []struct {
		num   []int
		count int
	}{
		{
			[]int{1,},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			8,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.num)
		count := countNodes(root)
		fmt.Println(count == test.count)
	}
}
