package linkedTag

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToArray(l *ListNode) []int {
	result := []int{}
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func CreateList(array []int) *ListNode {
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
