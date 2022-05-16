package treeTag

import "fmt"

/**
951. 翻转等价二叉树
我们可以为二叉树 T 定义一个翻转操作，如下所示：选择任意节点，然后交换它的左子树和右子树。

只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转等价于二叉树 Y。

编写一个判断两个二叉树是否是翻转等价的函数。这些树由根节点 root1 和 root2 给出。

示例：

输入：root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
输出：true
解释：We flipped at nodes with values 1, 3, and 5.

提示：

每棵树最多有 100 个节点。
每棵树中的每个值都是唯一的、在 [0, 99] 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flip-equivalent-binary-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

func TestflipEquiv() {
	tests := []struct {
		num1   []int
		num2   []int
		result bool
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			true,
		},
		{
			[]int{0, -1, 1},
			[]int{0, -1, 1},
			true,
		},
		{
			[]int{0, 3, 1, -1, -1, -1, 2},
			[]int{0, 3, 1, 2},
			false,
		},
		{
			[]int{0, 1, 3, -1, -1, -1, 2},
			[]int{0, 3, 1, 2},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 3, 2, 7, 6, 4, 5},
			true,
		},
	}

	for _, test := range tests {
		root1 := ArrayToTree(test.num1)
		root2 := ArrayToTree(test.num2)
		result := flipEquiv(root1, root2)
		fmt.Println(result == test.result)
	}
}
