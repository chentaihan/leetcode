package treeTag

/**
783. 二叉搜索树结点最小距离
给定一个二叉搜索树的根结点 root, 返回树中任意两节点的差的最小值。

示例：

输入: root = [4,2,6,1,3,null,null]
输出: 1
解释:
注意，root是树结点对象(TreeNode object)，而不是数组。

给定的树 [4,2,6,1,3,null,null] 可表示为下图:

          4
        /   \
      2      6
     / \
    1   3

最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
注意：

二叉树的大小范围在 2 到 100。
二叉树总是有效的，每个节点的值都是整数，且不重复。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：根节点和左子树的最大值比较，和右子树最小值比较
*/

func minDiffInBST(root *TreeNode) int {
	minValue := 1<<63 - 1
	_minDiffInBST(root, &minValue)
	return minValue
}

func _minDiffInBST(root *TreeNode, value *int) {
	if root != nil {
		if root.Left != nil {
			left := root.Val - getRightMax(root.Left)
			if left < *value {
				*value = left
			}
		}
		if root.Right != nil {
			left := getLeftMin(root.Right) - root.Val
			if left < *value {
				*value = left
			}
		}
		_minDiffInBST(root.Left, value)
		_minDiffInBST(root.Right, value)
	}
}

func getLeftMin(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func getRightMax(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func minDiffInBSTEx(root *TreeNode) int {
	var prev *TreeNode
	minValue := 1<<63 - 1
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root != nil {
			f(root.Left)
			if prev != nil {
				value := root.Val - prev.Val
				if value < minValue {
					minValue = value
				}
			}
			prev = root
			f(root.Right)
		}
	}
	f(root)
	return minValue
}
