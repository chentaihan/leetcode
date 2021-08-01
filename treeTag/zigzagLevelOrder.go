package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
103. 二叉树的锯齿形层次遍历
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

3
/ \
9  20
/  \
15   7
返回锯齿形层次遍历如下：

[
[3],
[20,9],
[15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	var f func(root *TreeNode, level int)
	f = func(root *TreeNode, level int) {
		if root != nil {
			if level >= len(ret) {
				ret = append(ret, []int{root.Val})
			} else {
				ret[level] = append(ret[level], root.Val)
			}
			f(root.Left, level+1)
			f(root.Right, level+1)
		}
	}
	f(root, 0)
	for i := 1; i < len(ret); i += 2 {
		list := ret[i]
		start, end := 0, len(list)-1
		for start < end {
			list[start], list[end] = list[end], list[start]
		}
	}
	return ret
}

func zigzagLevelOrderEx(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	type Item struct {
		node  *TreeNode
		level int
	}
	queue := []Item{{root, 0}}
	for len(queue) > 0 {
		item := queue[0]
		if item.level >= len(ret) {
			ret = append(ret, []int{item.node.Val})
		} else {
			ret[item.level] = append(ret[item.level], item.node.Val)
		}
		if item.node.Left != nil {
			queue = append(queue, Item{item.node.Left, item.level + 1})
		}
		if item.node.Right != nil {
			queue = append(queue, Item{item.node.Right, item.level + 1})
		}
		queue = queue[1:]
	}
	for i := 1; i < len(ret); i += 2 {
		list := ret[i]
		start, end := 0, len(list)-1
		for start < end {
			list[start], list[end] = list[end], list[start]
			start++
			end--
		}
	}
	return ret
}

func TestzigzagLevelOrder() {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{1, 2, 3, 4, -1, -1, 5},
			[][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		ret := zigzagLevelOrderEx(root)
		fmt.Println(common.IntEqualEx(ret, test.result))
	}

}
