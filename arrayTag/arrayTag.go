package arrayTag

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(array []int) *ListNode {
	var root, cur *ListNode
	for _, val := range array {
		newNode := &ListNode{
			Val: val,
		}
		if root == nil {
			root, cur = newNode, newNode
		} else {
			cur.Next = newNode
			cur = cur.Next
		}
	}
	return root
}

func nodeListToArray(root *ListNode) []int {
	result := []int{}
	for root != nil {
		result = append(result, root.Val)
		root = root.Next
	}
	return result
}

type point struct{
	x int
	y int
}
