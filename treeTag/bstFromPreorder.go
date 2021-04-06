package treeTag

/**
1008. 前序遍历构造二叉搜索树
返回与给定前序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。

(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，前序遍历首先显示节点 node 的值，然后遍历 node.left，接着遍历 node.right。）

题目保证，对于给定的测试用例，总能找到满足要求的二叉搜索树。

示例：

输入：[8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]

提示：

1 <= preorder.length <= 100
1 <= preorder[i] <= 10^8
preorder 中的值互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
解析：先根节点，再左子树，后右子树
*/

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{
		Val: preorder[0],
	}
	for i := 1; i < len(preorder); i++ {
		_bstFromPreorder(root, preorder[i])
	}
	return root
}

func _bstFromPreorder(root *TreeNode, num int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: num,
		}
	}
	if root.Val > num {
		root.Left = _bstFromPreorder(root.Left, num)
	} else {
		root.Right = _bstFromPreorder(root.Right, num)
	}
	return root
}
