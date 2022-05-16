package treeTag

import (
	"fmt"
)

/**
124. 二叉树中的最大路径和
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。



示例 1：


输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：


输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42


提示：

树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func maxPathSum(root *TreeNode) int {
	maxValue := -1 << 63
	maxPathSumSub(root, &maxValue)
	return maxValue
}

func maxPathSumSub(root *TreeNode, maxValue *int) int {
	if root == nil {
		return 0
	}
	leftValue := maxPathSumSub(root.Left, maxValue)
	if leftValue < 0 {
		leftValue = 0
	}
	rightValue := maxPathSumSub(root.Right, maxValue)
	if rightValue < 0 {
		rightValue = 0
	}
	val := leftValue + root.Val + rightValue
	if val > *maxValue {
		*maxValue = val
	}
	if rightValue > leftValue {
		return rightValue + root.Val
	}
	return leftValue + root.Val
}

func TestmaxPathSum() {
	tests := []struct {
		array []int
		max   int
	}{
		{
			[]int{9, 6, -3, -1, -1, -6, 2, -1, -1, 2, -1, -6, -6, -6},
			16,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			18,
		},
		{
			[]int{2, -1},
			2,
		},
		{
			[]int{3, -1, -2},
			3,
		},
		{
			[]int{1, -2, 3},
			4,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.array)
		max := maxPathSum(root)
		fmt.Println(max == test.max)
	}
}
