package treeTag

/**
面试题54. 二叉搜索树的第k大节点
给定一棵二叉搜索树，请找出其中第k大的节点。



示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4


限制：

1 ≤ k ≤ 二叉搜索树元素个数

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func kthLargest(root *TreeNode, k int) int {
	result := 0
	_kthLargest(root, &k, &result)
	return result
}

func _kthLargest(root *TreeNode, k, result *int) {
	if root != nil {
		_kthLargest(root.Right, k, result)
		*k--
		if *k == 0 {
			*result = root.Val
			return
		}
		_kthLargest(root.Left, k, result)
	}
}
