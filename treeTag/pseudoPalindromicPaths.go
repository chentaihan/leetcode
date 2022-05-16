package treeTag

import "fmt"

/**
1457. 二叉树中的伪回文路径
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。

请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。



示例 1：



输入：root = [2,3,1,3,1,null,1]
输出：2
解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
     在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1] 存在回文排列 [1,2,1] 。
示例 2：

输入：root = [2,1,1,1,3,null,null,null,null,null,1]
输出：1
解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
     这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。
示例 3：

输入：root = [9]
输出：1

提示：

给定二叉树的节点数目在 1 到 10^5 之间。
节点值在 1 到 9 之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func pseudoPalindromicPaths(root *TreeNode) int {
	buf := make([]int, 10)
	count := 0
	_pseudoPalindromicPaths(root, buf, &count)
	return count
}

func _pseudoPalindromicPaths(root *TreeNode, buf []int, count *int) {
	if root == nil {
		return
	}
	buf[root.Val]++
	if root.Left == root.Right && isPalindromic(buf) {
		*count++
		buf[root.Val]--
		return
	}
	_pseudoPalindromicPaths(root.Left, buf, count)
	_pseudoPalindromicPaths(root.Right, buf, count)
	buf[root.Val]--
}

func isPalindromic(buf []int) bool {
	oneCount := 0
	for i := 0; i < len(buf); i++ {
		if buf[i]%2 == 1 {
			oneCount++
		}
	}
	return oneCount <= 1
}

func TestpseudoPalindromicPaths() {
	tests := []struct {
		nums  []int
		count int
	}{
		{
			[]int{2, 3, 1, 3, 1, -1, 1},
			2,
		},
	}

	for index, test := range tests {
		root := ArrayToTree(test.nums)
		result := pseudoPalindromicPaths(root)
		if result != test.count {
			fmt.Println(index, ": ", result, test.count)
		}
	}
}
