package treeTag

//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func postorderTraversal(root *TreeNode) []int {
	var list []int
	postorderTraversalUtil(root, &list)
	return list
}

func postorderTraversalUtil(root *TreeNode, list *[]int) {
	if root != nil {
		postorderTraversalUtil(root.Left, list)
		postorderTraversalUtil(root.Right, list)
		*list = append(*list, root.Val)
	}
}
