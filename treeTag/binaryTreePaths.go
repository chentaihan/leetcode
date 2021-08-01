package treeTag

/**
257. 二叉树的所有路径
 */

import (
	"strconv"
	"strings"
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

func binaryTreePaths(root *TreeNode) []string {
	var rest []string
	var list []string
	var minPrint func(node *TreeNode)
	minPrint = func(node *TreeNode) {
		if node != nil {
			list = append(list, strconv.Itoa(node.Val))
			if node.Left == nil && node.Right == nil {
				data := make([]string, len(list))
				copy(data, list)
				rest = append(rest, strings.Join(data, "->"))
			}
			minPrint(node.Left)
			minPrint(node.Right)
			list = list[:len(list)-1]
		}
	}
	minPrint(root)
	return rest
}

func binaryTreePathsEx(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var list []string
	_binaryTreePathsEx(root, &list, "")
	return list
}

func _binaryTreePathsEx(root *TreeNode, list *[]string, path string) {
	if root != nil {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*list = append(*list, path)
		}
		if root.Left != nil {
			_binaryTreePathsEx(root.Left, list, path+"->")
		}
		if root.Right != nil {
			_binaryTreePathsEx(root.Right, list, path+"->")
		}
	}
}

func TestbinaryTreePaths(){
	tests := []struct {
		nums       []int
		result []string
	}{
		{
			[]int{1},
			[]string{"1"},
		},
		{
			[]int{1, 2},
			[]string{"1->2"},
		},
		{
			[]int{1, 2, 2},
			[]string{"1->2","1->2"},
		},
		{
			[]int{1, 2, 2, 3},
			[]string{"1->2->3","1->2"},
		},
		{
			[]int{1, 2, 2, 3, 3},
			[]string{"1->2->3","1->2->3","1->2"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3},
			[]string{"1->2->3","1->2->3","1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3},
			[]string{"1->2->3","1->2->3","1->2->3","1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4},
			[]string{"1->2->3->4","1->2->3","1->2->3","1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4},
			[]string{"1->2->3->4","1->2->3->4","1->2->3","1->2->3","1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, -1, -1, -1, 4},
			[]string{"1->2->3->4","1->2->3","1->2->3->4","1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4},
			[]string{"1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5},
			[]string{"1->2->3->4->5","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5},
			[]string{"1->2->3->4->5","1->2->3->4->5","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4","1->2->3->4"},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := binaryTreePathsEx(root)
		fmt.Println(common.StringEqual(test.result ,result) )
	}
}