package treeTag

import (
	"github.com/chentaihan/leetcode/common"
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
func AddTreeNode(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{
		Val: val,
	}
	if root == nil {
		return node
	}
	curNode := root
	for curNode != nil {
		if curNode.Val < val {
			if curNode.Right == nil {
				curNode.Right = node
				node.Parent = curNode
				break
			}
			curNode = curNode.Right
		} else {
			if curNode.Left == nil {
				curNode.Left = node
				node.Parent = curNode
				break
			}
			curNode = curNode.Left
		}
	}
	return root
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
	return nil
}

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
	if root == nil {
		return 0
	}
	return 1 + TreeNodeCount(root.Left) + TreeNodeCount(root.Right)
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
按照前序遍历将节点的值存入list中
*/
func DLRList(root *TreeNode, list *[]int) {
	if root != nil {
		*list = append(*list, root.Val)
		DLRList(root.Left, list)
		DLRList(root.Right, list)
	}
}

/**
按照前序遍历将节点的值存入map中,map的key的规则：root=1,left=2N,right=2N+1
*/
func DLRMap(root *TreeNode, m map[int]int, index int) {
	if root != nil {
		m[index] = root.Val
		DLRMap(root.Left, m, index*2)
		DLRMap(root.Right, m, index*2+1)
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
层序遍历
*/
func TreeFloor(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := Queue{}
	ret := []int{}
	queue.Push(root)
	for !queue.Empty() {
		node := queue.Pop()
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
		ret = append(ret, node.Val)
	}
	return ret
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
最大节点
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

/**
将数组转换成一个普通的二叉树
如果数组的值为-1，表示这个地方没有对应节点
*/
func ArrayToTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[0],
	}
	queue := QueueEx{}
	queue.Push(root, 0)
	for !queue.Empty() {
		node, index := queue.Pop()
		leftIndex := 2*index + 1
		if leftIndex < len(nums) && nums[leftIndex] != -1 {
			leftNode := &TreeNode{
				Val: nums[leftIndex],
			}
			node.Left = leftNode
			queue.Push(leftNode, leftIndex)
		}
		rightIndex := 2*index + 2
		if rightIndex < len(nums) && nums[rightIndex] != -1 {
			rightNode := &TreeNode{
				Val: nums[rightIndex],
			}
			node.Right = rightNode
			queue.Push(rightNode, rightIndex)
		}
	}
	return root
}

func TestArrayToTree() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{2, 3},
			[]int{2, 3},
		},
		{
			[]int{2, 3, 3},
			[]int{2, 3, 3},
		},
		{
			[]int{2, 3, 3, 4},
			[]int{2, 3, 3, 4},
		},
		{
			[]int{2, 3, 3, 4, 5},
			[]int{2, 3, 3, 4, 5},
		},
		{
			[]int{2, 3, 3, 4, 5, 4},
			[]int{2, 3, 3, 4, 5, 4},
		},
		{
			[]int{2, 3, 3, 4, 5, -1, 4},
			[]int{2, 3, 3, 4, 5, 4},
		},
		{
			[]int{2, 3, 3, 4, -1, 5, -1, 4},
			[]int{2, 3, 3, 4, 5, 4},
		},
		{
			[]int{2, 3, 3, 4, -1, 5, -1, 4},
			[]int{2, 3, 3, 4, 5, 4},
		},
		{
			[]int{2, 3, -1, 4, -1, -1, -1, 5, 4},
			[]int{2, 3, 4, 5, 4},
		},
		{
			[]int{2, 3, -1, 4, -1, -1, -1, 5, 4, -1, -1, -1, -1, -1, -1, 1, 2, 3, 4},
			[]int{2, 3, 4, 5, 4, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.nums)
		result := TreeFloor(root)
		fmt.Println(common.IntEqual(result, test.result))
	}
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
