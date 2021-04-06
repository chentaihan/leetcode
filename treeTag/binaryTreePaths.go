package treeTag

/**
257. 二叉树的所有路径
*/

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"strconv"
	"strings"
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
	var rest []string
	var minPrint func(node *TreeNode, path string)
	minPrint = func(node *TreeNode, path string) {
		if node != nil {
			path += strconv.Itoa(node.Val)
			if node.Left == nil && node.Right == nil {
				rest = append(rest, path)
				return
			}
			minPrint(node.Left, path+"->")
			minPrint(node.Right, path+"->")
		}
	}
	minPrint(root, "")
	return rest
}

func TestbinaryTreePaths() {
	tests := []struct {
		nums   []int
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
			[]string{"1->2", "1->2"},
		},
		{
			[]int{1, 2, 2, 3},
			[]string{"1->2->3", "1->2"},
		},
		{
			[]int{1, 2, 2, 3, 3},
			[]string{"1->2->3", "1->2->3", "1->2"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3},
			[]string{"1->2->3", "1->2->3", "1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3},
			[]string{"1->2->3", "1->2->3", "1->2->3", "1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4},
			[]string{"1->2->3->4", "1->2->3", "1->2->3", "1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4},
			[]string{"1->2->3->4", "1->2->3->4", "1->2->3", "1->2->3", "1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, -1, -1, -1, 4},
			[]string{"1->2->3->4", "1->2->3", "1->2->3->4", "1->2->3"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4},
			[]string{"1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5},
			[]string{"1->2->3->4->5", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4"},
		},
		{
			[]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5},
			[]string{"1->2->3->4->5", "1->2->3->4->5", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4", "1->2->3->4"},
		},
	}
	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := binaryTreePathsEx(root)
		fmt.Println(common.StringEqual(test.result, result))
	}
}
