package treeTag

/**
1305. 两棵二叉搜索树中的所有元素
给你 root1 和 root2 这两棵二叉搜索树。

请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.

 

示例 1：



输入：root1 = [2,1,4], root2 = [1,0,3]
输出：[0,1,1,2,3,4]
示例 2：

输入：root1 = [0,-10,10], root2 = [5,1,7,0,2]
输出：[-10,0,0,1,2,5,7,10]
示例 3：

输入：root1 = [], root2 = [5,1,7,0,2]
输出：[0,1,2,5,7]
示例 4：

输入：root1 = [0,-10,10], root2 = []
输出：[-10,0,10]
示例 5：

输入：root1 = [1,null,8], root2 = [8,1]
输出：[1,1,8,8]
 
提示：

每棵树最多有 5000 个节点。
每个节点的值在 [-10^5, 10^5] 之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var list1, list2 []int
	getNodeList(root1, &list1)
	getNodeList(root2, &list2)
	res := make([]int, len(list1)+len(list2))
	index, index1, index2 := 0, 0, 0
	for index1 < len(list1) && index2 < len(list2) {
		if list1[index1] < list2[index2] {
			res[index] = list1[index1]
			index1++
		} else {
			res[index] = list2[index2]
			index2++
		}
		index++
	}
	for index1 < len(list1) {
		res[index] = list1[index1]
		index1++
		index++
	}
	for index2 < len(list2) {
		res[index] = list2[index2]
		index2++
		index++
	}
	return res
}

func getNodeList(root *TreeNode, list *[]int) {
	if root != nil {
		getNodeList(root.Left, list)
		*list = append(*list, root.Val)
		getNodeList(root.Right, list)
	}
}
