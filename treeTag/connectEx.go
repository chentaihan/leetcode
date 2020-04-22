package treeTag

/**
117. 填充每个节点的下一个右侧节点指针 II
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func connectEx(root *Node) *Node {
	if root == nil {
		return root
	}
	type item struct {
		node  *Node
		level int
	}
	queue := make([]*item, 0, 1024)
	queue = append(queue, &item{root, 0})
	var prev *item
	for len(queue) > 0 {
		cur := queue[0]
		if cur.node.Left != nil {
			queue = append(queue, &item{cur.node.Left, cur.level + 1})
		}
		if cur.node.Right != nil {
			queue = append(queue, &item{cur.node.Right, cur.level + 1})
		}
		if prev != nil && prev.level == cur.level {
			prev.node.Next = cur.node
		}
		prev = cur
		queue = queue[1:]
	}
	return root
}
