package treeTag

/**
559. N叉树的最大深度
给定一个 N 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

例如，给定一个 3叉树 :https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png

我们应返回其最大深度，3。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type NodeN struct {
	Val      int
	Children []*NodeN
}

func maxDepthN(root *NodeN) int {
	if root == nil {
		return 0
	}
	maxLevel := 1
	_maxDepthN(root, 1, &maxLevel)
	return maxLevel
}

func _maxDepthN(root *NodeN, level int, maxLevel *int) {
	if root != nil {
		if level > *maxLevel {
			*maxLevel = level
		}
		for i := 0; i < len(root.Children); i++ {
			_maxDepthN(root.Children[i], level+1, maxLevel)
		}
	}
}

func maxDepthNEx(root *NodeN) int {
	if root == nil {
		return 0
	}
	maxValue := 0
	for i := 0; i < len(root.Children); i++ {
		value := maxDepthNEx(root.Children[i])
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue + 1
}
