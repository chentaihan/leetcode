package treeTag

/**
1038. 从二叉搜索树到更大和树
给出二叉 搜索 树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。


示例：https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/05/03/tree.png

输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

提示：

树中的节点数介于 1 和 100 之间。
每个节点的值介于 0 和 100 之间。
给定的树为二叉搜索树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	_bstToGst(root, &sum)
	return root
}

func _bstToGst(root *TreeNode, sum *int) {
	if root != nil {
		_bstToGst(root.Right, sum)
		val := root.Val
		root.Val += *sum
		*sum += val
		_bstToGst(root.Left, sum)
	}
}
