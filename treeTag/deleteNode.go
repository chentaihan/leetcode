package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-node-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	target, parent := findTreeNode(root, key)
	if target == nil {
		return root
	}
	if target.Left != nil {
		cur, parent := target.Left, target.Left
		for cur.Right != nil {
			parent = cur
			cur = cur.Right
		}
		target.Val = cur.Val
		if cur == target.Left {
			target.Left = cur.Left
		} else {
			parent.Right = nil
		}
		return root
	}
	if target.Right != nil {
		cur, parent := target.Right, target.Right
		for cur.Left != nil {
			parent = cur
			cur = cur.Left
		}
		target.Val = cur.Val
		if cur == target.Right {
			target.Right = cur.Right
		} else {
			parent.Left = nil
		}
		return root
	}
	if parent != nil {
		if parent.Left == target {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else {
		return nil
	}
	return root
}

func findTreeNode(root *TreeNode, key int) (*TreeNode, *TreeNode) {
	if root.Val == key {
		return root, nil
	}
	cur, parent := root, root
	for cur != nil {
		if cur.Val == key {
			return cur, parent
		}
		parent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return nil, nil
}

func TestdeleteNode() {
	tests := []struct {
		nums   []int
		key    int
		result []int
	}{
		//{
		//	[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
		//	5,
		//	[]int{10, 3, 20, 2, 7, 15, 30, 1, 6, 8, 12, 18, 25, 40},
		//},
		//{
		//	[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
		//	20,
		//	[]int{10, 5, 18, 2, 7, 15, 30, 1, 3, 6, 8, 12, 25, 40},
		//},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			1,
			[]int{10, 5, 20, 2, 7, 15, 30, 3, 6, 8, 12, 18, 25, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			3,
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 6, 8, 12, 18, 25, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			6,
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 8, 12, 18, 25, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			10,
			[]int{8, 5, 20, 2, 7, 15, 30, 1, 3, 6, 12, 18, 25, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			8,
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 12, 18, 25, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			25,
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			40,
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			2,
			[]int{10, 5, 20, 1, 7, 15, 30, 3, 6, 8, 12, 18, 25, 40},
		},
		{
			[]int{10, 5, 20, 2, 7, 15, 30, 1, 3, 6, 8, 12, 18, 25, 40},
			7,
			[]int{10, 5, 20, 2, 6, 15, 30, 1, 3, 8, 12, 18, 25, 40},
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		root = deleteNode(root, test.key)
		result := TreeToArray(root)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
