package treeTag

/**
2331. 计算布尔二叉树的值
给你一棵 完整二叉树的根，这棵树有以下特征：

叶子节点要么值为0要么值为1，其中0 表示False，1 表示True。
非叶子节点 要么值为 2要么值为 3，其中2表示逻辑或OR ，3表示逻辑与AND。
计算一个节点的值方式如下：

如果节点是个叶子节点，那么节点的 值为它本身，即True或者False。
否则，计算两个孩子的节点值，然后将该节点的运算符对两个孩子值进行 运算。
返回根节点root的布尔运算值。

完整二叉树是每个节点有 0个或者 2个孩子的二叉树。

叶子节点是没有孩子的节点。



示例 1：



输入：root = [2,1,3,null,null,0,1]
输出：true
解释：上图展示了计算过程。
AND 与运算节点的值为 False AND True = False 。
OR 运算节点的值为 True OR False = True 。
根节点的值为 True ，所以我们返回 true 。
示例 2：

输入：root = [0]
输出：false
解释：根节点是叶子节点，且值为 false，所以我们返回 false 。


提示：

树中节点数目在[1, 1000]之间。
0 <= Node.val <= 3
每个节点的孩子数为0 或2。
叶子节点的值为0或1。
非叶子节点的值为2或3 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/evaluate-boolean-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		}
		return true
	}
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == 2 {
		return left || right
	}
	return left && right
}
