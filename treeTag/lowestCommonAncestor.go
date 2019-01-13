package treeTag

import "fmt"

/**
235. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。
                  100
           50                200
      25       75        150          250
   10    35  60  90   120   180    220    300

思路1
1：分别查找两个节点，记录从根节点到目标节点的路径
2：遍历两个路径，第一个不一样的节点的上一个节点即为所求
我去，打败0%的用户，思路不对

思路2
1：两个节点的父节点一定在两个节点之间，必定存在：(root.Val - q.Val)*(root.Val - p.Val) <= 0
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	left := _getTreeNodePath(root, p)
	right := _getTreeNodePath(root, q)
	l := len(left)
	if len(right) < l {
		l = len(right)
	}
	i := 0
	for ; i < l; i++ {
		if left[i].Val != right[i].Val {
			break
		}
	}
	return left[i-1]
}

func lowestCommonAncestorEx(root, p, q *TreeNode) *TreeNode {
	if (root.Val - q.Val)*(root.Val - p.Val) <= 0 {
		return root
	}
	if root.Val < p.Val {
		return lowestCommonAncestorEx(root.Right, p, q)
	}
	return lowestCommonAncestorEx(root.Left, p, q)
}

func _getTreeNodePath(root, p *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	for root != nil {
		if root.Val == p.Val {
			list = append(list, root)
			break
		}
		list = append(list, root)
		if root.Val < p.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return list
}

func TestLowestCommonAncestor() {
	tests := []struct {
		nums []int
		p    int
		q    int
		ret  int
	}{
		{
			[]int{2, 1, 3},
			1,
			2,
			2,
		},
		{
			[]int{2, 1, 3},
			1,
			3,
			2,
		},
		{
			[]int{2, 1, 4, -1, -1, 3},
			4,
			3,
			4,
		},
		{
			[]int{2, 1, 4, -1, -1, 3},
			1,
			3,
			2,
		},
		{
			[]int{100, 50, 200, 25, 75, 150, 250, 10, 35, 60, 90, 120, 180, 220, 3003},
			220,
			300,
			250,
		},
		{
			[]int{100, 50, 200, 25, 75, 150, 250, 10, 35, 60, 90, 120, 180, 220, 3003},
			250,
			120,
			200,
		},
		{
			[]int{100, 50, 200, 25, 75, 150, 250, 10, 35, 60, 90, 120, 180, 220, 3003},
			10,
			120,
			100,
		},
		{
			[]int{100, 50, 200, 25, 75, 150, 250, 10, 35, 60, 90, 120, 180, 220, 3003},
			35,
			60,
			50,
		},
		{
			[]int{100, 50, 200, 25, 75, 150, 250, 10, 35, 60, 90, 120, 180, 220, 3003},
			10,
			300,
			100,
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		p := &TreeNode{Val: test.p}
		q := &TreeNode{Val: test.q}
		ret := lowestCommonAncestorEx(root, p, q)
		fmt.Println(ret.Val == test.ret)
	}
}
