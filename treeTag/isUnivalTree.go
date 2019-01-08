package treeTag

/**
965. 单值二叉树
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
只有给定的树是单值二叉树时，才返回 true；否则返回 false。
 */

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return _isUnivalTree(root, root.Val)
}

func _isUnivalTree(root *TreeNode, val int) bool {
	if root != nil {
		if root.Val != val {
			return false
		}
		return _isUnivalTree(root.Left, root.Val) && _isUnivalTree(root.Right, root.Val)
	}
	return true
}