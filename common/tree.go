package common

import (
	"fmt"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

/**
构造一棵树
*/
func CreateTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		AddTreeNode(root, nums[i])
	}
	return root
}

/**
树添加节点
*/
func AddTreeNode(root *TreeNode, val int) {
	node := &TreeNode{
		Val: val,
	}
	for root != nil {
		if root.Val < val {
			if root.Right == nil {
				root.Right = node
				node.Parent = root
				break
			}
			root = root.Right
		} else {
			if root.Left == nil {
				root.Left = node
				node.Parent = root
				break
			}
			root = root.Left
		}
	}
}

/**
查找节点
*/
func SearchTreeNode(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil}

/**
删除节点
rootNode用双指针是为了能删除根节点
 */
func DeleteTreeNode(rootNode **TreeNode, val int) bool {
	root := *rootNode
	targetNode := SearchTreeNode(root, val)
	if targetNode == nil {
		return false
	}
	//删除不了跟节点
	if targetNode == root {
		*rootNode = nil
		return true
	}
	var removeNode *TreeNode
	if targetNode.Left != nil || targetNode.Right != nil {
		if targetNode.Right != nil {
			//右子树最小值代替当节点
			removeNode = MinNode(targetNode)
			targetNode.Val = removeNode.Val
			//删除右子树最小值对应的节点
			parentNode := removeNode.Parent
			if parentNode.Left == removeNode {
				parentNode.Left = nil
			} else {
				parentNode.Right = nil
			}
		} else {
			//左子树最大值，即左子树的根节点代替当前节点
			removeNode = MaxNode(targetNode)
			parentNode := targetNode.Parent
			parentNode.Left = targetNode.Left
			targetNode.Left.Parent = parentNode
		}
	} else {
		//删除代替要被删除节点的叶子节点
		parentNode := targetNode.Parent
		if parentNode.Left == targetNode {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
	}

	return true
}

/**
计算树的深度
*/
func TreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := TreeDepth(root.Left) + 1
	rightDepth := TreeDepth(root.Right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

/**
计算树总节点数
*/
func TreeNodeCount(root *TreeNode) int {
	count := 0
	treeNodeCount(root, &count)
	return count
}

func treeNodeCount(root *TreeNode, count *int) {
	if root != nil {
		*count++
		treeNodeCount(root.Left, count)
		treeNodeCount(root.Right, count)
	}
}

/**
前序遍历
*/
func DLR(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val)
		DLR(root.Left)
		DLR(root.Right)
	}
}

/**
中序遍历
*/
func LDR(root *TreeNode) {
	if root != nil {
		LDR(root.Left)
		fmt.Print(root.Val)
		LDR(root.Right)
	}
}

/**
后序遍历
*/
func LRD(root *TreeNode) {
	if root != nil {
		LRD(root.Left)
		LRD(root.Right)
		fmt.Print(root.Val)
	}
}

/**
最小节点
*/
func MinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

/**
最小节点
*/
func MaxNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func TestTree() {
	tests := []struct {
		nums       []int
		depth      int
		count      int
		minVal     int
		maxVal     int
		findVal    int
		findResult bool
	}{
		{
			[]int{5, 4, 8, 2, 3, 9},
			4,
			6,
			2,
			9,
			5,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 7},
			4,
			7,
			2,
			9,
			4,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1},
			4,
			7,
			1,
			9,
			9,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6},
			4,
			8,
			1,
			9,
			1,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7},
			4,
			9,
			1,
			9,
			10,
			false,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15},
			4,
			10,
			1,
			15,
			15,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12},
			5,
			11,
			1,
			15,
			3,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20},
			5,
			12,
			1,
			20,
			10,
			false,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14},
			6,
			13,
			1,
			20,
			0,
			false,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11},
			6,
			14,
			1,
			20,
			13,
			false,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25},
			6,
			15,
			1,
			25,
			12,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18},
			6,
			16,
			1,
			25,
			18,
			true,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18, 30},
			7,
			17,
			1,
			30,
			13,
			false,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		fmt.Println(TreeDepth(root) == test.depth)
		count := TreeNodeCount(root)
		fmt.Println(count == test.count)
		fmt.Println(test.minVal == MinNode(root).Val)
		fmt.Println(test.maxVal == MaxNode(root).Val)
		fmt.Println(test.findResult == (SearchTreeNode(root, test.findVal) != nil))
	}
}

func TestDeleteTreeNode() {
	tests := []struct {
		nums   []int
		delVal int
		count  int
	}{
		{
			[]int{5},
			5,
			0,
		},
		{
			[]int{5, 4, 8, 2, 3, 9},
			4,
			5,
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 7},
			4,
			6,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30},
			10,
			14,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30},
			5,
			14,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30},
			30,
			14,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30},
			22,
			14,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30},
			20,
			14,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30},
			17,
			14,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30, 29, 27, 26, 28},
			27,
			18,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30, 29, 27, 26, 28},
			29,
			18,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30, 29, 27, 26, 28},
			26,
			18,
		},
		{
			[]int{15, 10, 20, 5, 13, 17, 25, 2, 6, 11, 14, 16, 19, 22, 30, 29, 27, 26, 28},
			28,
			18,
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		DeleteTreeNode(&root, test.delVal)
		fmt.Println(TreeNodeCount(root) == test.count)
	}
}
