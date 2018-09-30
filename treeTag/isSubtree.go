package treeTag

import "fmt"

/**
572. 另一个树的子树
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
 */

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return s != nil && (isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t))
}

func TestisSubTree() {
	tests := []struct {
		nums1  []int
		nums2  []int
		result bool
	}{
		{
			[]int{3, 4, 5, 1, 2, -1, -1, 0},
			[]int{4, 1, 2},
			false,
		},
		{
			[]int{},
			[]int{},
			false,
		},
		{
			[]int{10},
			[]int{10},
			true,
		},
		{
			[]int{10, 3},
			[]int{10, 3},
			true,
		},
		{
			[]int{10, 3, 12},
			[]int{10, 3, 12},
			true,
		},
		{
			[]int{10, 3, 12},
			[]int{10, 3},
			false,
		},
		{
			[]int{10, 3, 12},
			[]int{10, -1, 12},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{2, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{3, 6, 7},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{2, 4, 5, 8},
			true,
		},
	}
	for _, test := range tests {
		t1 := ArrayToTree(test.nums1)
		t2 := ArrayToTree(test.nums2)
		result := isSubtree(t1, t2)
		fmt.Println(result == test.result)
	}
}
