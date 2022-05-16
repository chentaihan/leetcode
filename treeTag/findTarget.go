package treeTag

import "fmt"

/**
653. 两数之和 IV - 输入 BST
给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

案例 1:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

输出: True


案例 2:

输入:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findTarget(root *TreeNode, k int) bool {
	return findTarget1(root, root, k)
}

func findTarget1(root, parent *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	value := k - root.Val
	if value < root.Val {
		if findTargetFind(parent.Left, value, root.Val) {
			return true
		}
	} else {
		if findTargetFind(parent.Right, value, root.Val) {
			return true
		}
	}
	return findTarget1(root.Left, parent, k) || findTarget1(root.Right, parent, k)
}

func findTargetFind(root *TreeNode, first, second int) bool {
	if root == nil {
		return false
	}
	if root.Val == first {
		if first != second {
			return true
		}
		return (root.Left != nil && root.Val == root.Left.Val) || (root.Right != nil && root.Val == root.Right.Val)
	}
	if root.Val > first {
		return findTargetFind(root.Left, first, second)
	}
	return findTargetFind(root.Right, first, second)
}

//-----------------------------方法二 -----------------------------------------------------
func findTargetEx(root *TreeNode, k int) bool {
	m := make(map[int]bool, 0)
	return findTargetEx1(root, k, m)
}

func findTargetEx1(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}
	if _findTargetEx(root, k, m) {
		return true
	}
	return findTargetEx1(root.Left, k, m) || findTargetEx1(root.Right, k, m)
}

func _findTargetEx(root *TreeNode, k int, m map[int]bool) bool {
	if root != nil {
		value := k - root.Val
		if m[value] && m[root.Val] {
			if value == root.Val {
				return (root.Left != nil && root.Val == root.Left.Val) || (root.Right != nil && root.Val == root.Right.Val)
			}
			return true
		}
		m[value] = true
		return _findTargetEx(root.Left, k, m) || _findTargetEx(root.Right, k, m)
	}
	return false
}

func TestfindTarget() {
	tests := []struct {
		nums   []int
		k      int
		result bool
	}{
		{
			[]int{2, -1, 3},
			6,
			false,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			28,
			false,
		},
		{
			[]int{2, -1, 3},
			5,
			true,
		},
		{
			[]int{0, -2, 2, -3, -1, -1, 4},
			-5,
			true,
		},
		{
			[]int{2, 1, 3},
			4,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			9,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			11,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			13,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			5,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			6,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			7,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			8,
			true,
		},
		{
			[]int{5, 3, 6, 2, 4, -1, 7},
			12,
			true,
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := findTargetEx(root, test.k)
		fmt.Println(result == test.result)
	}
}
