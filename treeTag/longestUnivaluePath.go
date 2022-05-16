package treeTag

import "fmt"

/**
687. 最长同值路径
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:

2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:

2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-univalue-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxCount := 0
	_longestUnivaluePath(root, root.Val, &maxCount)
	return maxCount
}

func _longestUnivaluePath(root *TreeNode, value int, maxValue *int) int {
	if root == nil {
		return 0
	}
	left := _longestUnivaluePath(root.Left, root.Val, maxValue)
	right := _longestUnivaluePath(root.Right, root.Val, maxValue)
	if left+right > *maxValue {
		*maxValue = left + right
	}
	if root.Val == value {
		return max(left, right) + 1
	}
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestlongestUnivaluePath() {
	tests := []struct {
		nums  []int
		count int
	}{
		{
			[]int{5, 4, 5, 1, 1, -1, 5},
			2,
		},
		{
			[]int{1, 4, 5, 4, 4, -1, 5},
			2,
		},
		{
			[]int{1, 2, 2, 2, 2, 2, 2, 2},
			3,
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		count := longestUnivaluePath(root)
		fmt.Println(count == test.count)
	}
}
