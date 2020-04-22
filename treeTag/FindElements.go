package treeTag

/**
1261. 在受污染的二叉树中查找元素
给出一个满足下述规则的二叉树：

root.val == 0
如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
现在这个二叉树受到「污染」，所有的 treeNode.val 都变成了 -1。

请你先还原二叉树，然后实现 FindElements 类：

FindElements(TreeNode* root) 用受污染的二叉树初始化对象，你需要先把它还原。
bool find(int target) 判断目标值 target 是否存在于还原后的二叉树中并返回结果。
 

示例 1：



输入：
["FindElements","find","find"]
[[[-1,null,-1]],[1],[2]]
输出：
[null,false,true]
解释：
FindElements findElements = new FindElements([-1,null,-1]);
findElements.find(1); // return False
findElements.find(2); // return True
示例 2：



输入：
["FindElements","find","find","find"]
[[[-1,-1,-1,-1,-1]],[1],[3],[5]]
输出：
[null,true,true,false]
解释：
FindElements findElements = new FindElements([-1,-1,-1,-1,-1]);
findElements.find(1); // return True
findElements.find(3); // return True
findElements.find(5); // return False
示例 3：



输入：
["FindElements","find","find","find","find"]
[[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
输出：
[null,true,false,false,true]
解释：
FindElements findElements = new FindElements([-1,null,-1,-1,null,-1]);
findElements.find(2); // return True
findElements.find(3); // return False
findElements.find(4); // return False
findElements.find(5); // return True
 

提示：

TreeNode.val == -1
二叉树的高度不超过 20
节点的总数在 [1, 10^4] 之间
调用 find() 的总次数在 [1, 10^4] 之间
0 <= target <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-elements-in-a-contaminated-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{
		root: root,
	}
	if root == nil {
		return fe
	}
	findElements(root, root)
	return fe
}

func findElements(root, parent *TreeNode) {
	if root != nil {
		if parent.Left == root {
			root.Val = 2*parent.Val + 1
		} else {
			root.Val = 2*parent.Val + 2
		}
		findElements(root.Left, root)
		findElements(root.Right, root)
	}
}

func (this *FindElements) Find(targetValue int) bool {
	var ops []int
	target := targetValue
	for target > 0 {
		parent := (target - 1) / 2
		op := 0
		if parent*2+2 == target {
			op = 1
		}
		target = parent
		ops = append(ops, op)
	}
	node := this.root
	for i := len(ops) - 1; i >= 0; i-- {
		if node == nil {
			return false
		}
		if ops[i] == 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if node == nil {
		return false
	}
	return node.Val == targetValue
}
