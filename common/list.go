package common

type ListNode struct {
	Val  int
	Next *ListNode
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

func NodeListToArray(root *ListNode) []int {
	result := []int{}
	for root != nil {
		result = append(result, root.Val)
		root = root.Next
	}
	return result
}

func GetListNodeCount(root *ListNode) int {
	count := 0
	for root != nil{
		count ++
		root = root.Next
	}
	return count
}