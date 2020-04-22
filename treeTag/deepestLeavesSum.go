package treeTag

/**
1302. 层数最深叶子节点的和
给你一棵二叉树，请你返回层数最深的叶子节点的和。

示例：

输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/deepest-leaves-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deepestLeavesSum(root *TreeNode) int {
	maxLevel, sum := 0, 0
	_deepestLeavesSum(root, 0, &maxLevel, &sum)
	return sum
}

func _deepestLeavesSum(root *TreeNode, level int, maxLevel, sum *int) {
	if root != nil {
		if level > *maxLevel {
			*maxLevel = level
			*sum = root.Val
		} else if level == *maxLevel {
			*sum += root.Val
		}
		_deepestLeavesSum(root.Left, level+1, maxLevel, sum)
		_deepestLeavesSum(root.Right, level+1, maxLevel, sum)
	}
}
