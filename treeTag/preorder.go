package treeTag

/**
589. N叉树的前序遍历
给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png

返回其后序遍历: [5,6,3,2,4,1].
*/

func preorder(root *NodeN) []int {
	var list []int
	_preorder(root, &list)
	return list
}

func _preorder(root *NodeN, list *[]int) {
	if root != nil {
		*list = append(*list, root.Val)
		for i := 0; i < len(root.Children); i++ {
			_preorder(root.Children[i], list)
		}
	}
}
