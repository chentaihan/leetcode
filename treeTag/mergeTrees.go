package treeTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
617. 合并二叉树
输入:
	Tree 1                     Tree 2
          10                        20
         / \                       / \
        3   12                     2   23
       /                           \   \
      2                             4   27
输出:
合并后的树:
	     30
	    / \
	   5   35
	  / \   \
	 2   4   27
*/

/**
第一版，思路如下，和好复杂，直接超时
N从1开始
left = 2N
right = 2N + 1
将tree 2的值存入数组array，遍历数组array，将array中的值存入tree 1中，存在的累加，不存在的创建节点
*/
func dlrMap(root *TreeNode, m map[int]int, index int) {
	if root != nil {
		m[index] = root.Val
		dlrMap(root.Left, m, index*2)
		dlrMap(root.Right, m, index*2+1)
	}
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	m := map[int]int{}
	dlrMap(t2, m, 1)
	dlrTree(t1, m, 1)
	return t1
}

func dlrTree(root *TreeNode, m map[int]int, index int) {
	if root != nil {
		//t1 和 t2存在相同节点则叠加
		if val, exist := m[index]; exist {
			root.Val += val
		}
		//t1不存在，t2存在，则在t1对应的位置创建节点
		if root.Left == nil {
			if _, exist := m[index*2]; exist {
				root.Left = &TreeNode{
					Val: 0,
				}
			}
		}
		dlrTree(root.Left, m, index*2)
		if root.Right == nil {
			if _, exist := m[index*2+1]; exist {
				root.Right = &TreeNode{
					Val: 0,
				}
			}
		}
		dlrTree(root.Right, m, index*2+1)
	}
}

/**
第二版，打败40%
*/
func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t1 != nil {
		if t2 != nil {
			t1.Val += t2.Val
		}

		var leftNode *TreeNode
		if t2 != nil {
			leftNode = t2.Left
		}
		if t1.Left == nil && leftNode != nil {
			t1.Left = &TreeNode{}
		}
		mergeTrees1(t1.Left, leftNode)

		var rightNode *TreeNode
		if t2 != nil {
			rightNode = t2.Right
		}
		if t1.Right == nil && rightNode != nil {
			t1.Right = &TreeNode{}
		}
		mergeTrees1(t1.Right, rightNode)
	}
	return t1
}

/**
第三版，打败80%
*/
func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	if t1.Left == nil && t2.Left != nil {
		t1.Left = t2.Left
	} else {
		mergeTrees2(t1.Left, t2.Left)
	}

	if t1.Right == nil && t2.Right != nil {
		t1.Right = t2.Right
	} else {
		mergeTrees2(t1.Right, t2.Right)
	}

	return t1
}

/**
膜拜版本
*/
func mergeTrees3(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t2 == nil {
		return t1
	} else if t1 == nil {
		return t2
	} else {
		t1.Val += t2.Val
	}
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func TestmergeTrees() {
	tests := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{
		{
			[]int{10, 3, 12, 2},
			[]int{20, 2, 23, 4, 27},
			[]int{30, 5, 2, 4, 35, 27},
		},
	}
	for _, test := range tests {
		t1 := CreateTree(test.nums1)
		t2 := CreateTree(test.nums2)
		mergeTrees1(t1, t2)
		list := []int{}
		DLRList(t1, &list)
		fmt.Println(common.IntEqual(test.result, list))
	}
}
