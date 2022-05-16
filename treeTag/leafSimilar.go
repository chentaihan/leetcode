package treeTag

/**
872. 叶子相似的树
请考虑一颗二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。

举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。

如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。

如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。

提示：

给定的两颗树可能会有 1 到 100 个结点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/leaf-similar-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：前序遍历拿到每棵树的叶子节点列表，对比两个链表
*/

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var list1 []int
	var list2 []int
	getleafs(root1, &list1)
	getleafs(root2, &list2)
	if len(list1) == len(list2) {
		for i := 0; i < len(list1); i++ {
			if list1[i] != list2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func getleafs(root *TreeNode, list *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*list = append(*list, root.Val)
		} else {
			getleafs(root.Left, list)
			getleafs(root.Right, list)
		}
	}
}
