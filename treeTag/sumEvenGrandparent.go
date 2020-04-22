package treeTag

/**
1315. 祖父节点值为偶数的节点和
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：

该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
如果不存在祖父节点值为偶数的节点，那么返回 0 。

示例：https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/10/1473_ex1.png

输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：18
解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sumEvenGrandparent(root *TreeNode) int {
	sum := 0
	_sumEvenGrandparent(root, nil, nil, &sum)
	return sum
}

func _sumEvenGrandparent(root, parent, pp *TreeNode, sum *int) {
	if root != nil {
		if pp != nil && pp.Val%2 == 0 {
			*sum += root.Val
		}
		_sumEvenGrandparent(root.Left, root, parent, sum)
		_sumEvenGrandparent(root.Right, root, parent, sum)
	}
}
