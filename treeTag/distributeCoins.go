package treeTag

import "fmt"

/**
979. 在二叉树中分配硬币 TODO
给定一个有 N 个结点的二叉树的根结点 root，树中的每个结点上都对应有 node.val 枚硬币，并且总共有 N 枚硬币。

在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。(移动可以是从父结点到子结点，或者从子结点移动到父结点。)。

返回使每个结点上只有一枚硬币所需的移动次数。



示例 1：https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/01/19/tree1.png



输入：[3,0,0]
输出：2
解释：从树的根结点开始，我们将一枚硬币移到它的左子结点上，一枚硬币移到它的右子结点上。
示例 2：



输入：[0,3,0]
输出：3
解释：从根结点的左子结点开始，我们将两枚硬币移到根结点上 [移动两次]。然后，我们把一枚硬币从根结点移到右子结点上。
示例 3：



输入：[1,0,2]
输出：2
示例 4：



输入：[1,0,0,null,3]
输出：4


提示：

1<= N <= 100
0 <= node.val <= N

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distribute-coins-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func distributeCoins(root *TreeNode) int {
	move, sum := 0, 0
	_distributeCoins(root, &move, &sum)
	return sum
}

func _distributeCoins(root *TreeNode, move, sum *int) {
	if root != nil {
		_distributeCoins(root.Left, move, sum)
		root.Val += *move
		if root.Val == 0 {
			*sum += 1
		} else if root.Val > 1 {
			*sum += root.Val - 1
		}
		*move = root.Val - 1
		_distributeCoins(root.Right, move, sum)
	}
}

func TestdistributeCoins() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 0, 0, -1, 3},
			4,
		},
		{
			[]int{3, 0, 0},
			2,
		},
		{
			[]int{0, 3, 0},
			3,
		},
		{
			[]int{1, 0, 2},
			2,
		},
	}
	for index, test := range tests {
		root := ArrayToTree(test.nums)
		result := distributeCoins(root)
		if result == test.result {
			fmt.Println(index, true)
		} else {
			fmt.Println(index, test.result, result)
		}
	}
}
