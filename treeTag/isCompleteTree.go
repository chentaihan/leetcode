package treeTag

import "fmt"

/**
958. 二叉树的完全性检验
给定一个二叉树，确定它是否是一个完全二叉树。

百度百科中对完全二叉树的定义如下：

若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）

示例 1：

输入：[1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
示例 2：

输入：[1,2,3,4,5,null,7]
输出：false
解释：值为 7 的结点没有尽可能靠向左侧。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	h := 0
	node := root
	for node != nil {
		node = node.Left
		h++
	}
	flag := 0
	return _isCompleteTree(root, h, 0, &flag)
}

func _isCompleteTree(root *TreeNode, maxFloor, curFloor int, flag *int) bool {
	if root != nil {
		return _isCompleteTree(root.Left, maxFloor, curFloor+1, flag) && _isCompleteTree(root.Right, maxFloor, curFloor+1, flag)
	}
	if curFloor < maxFloor-1 || curFloor > maxFloor {
		return false
	}
	if curFloor == maxFloor {
		if *flag == 1 {
			return false
		}
		return true
	}
	if *flag == 0 {
		*flag = 1
	}
	return true
}

func TestisCompleteTree() {
	tests := []struct {
		nums   []int
		result bool
	}{
		{
			[]int{1},
			true,
		},
		{
			[]int{1, 2},
			true,
		},
		{
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			true,
		},
		{
			[]int{1, 2, 3, -1, 4},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, -1, 5},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, -1, 6},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			true,
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := isCompleteTree(root)
		fmt.Println(result == test.result)
	}
}
