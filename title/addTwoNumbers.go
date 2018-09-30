package title

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(val string) *ListNode {
	var retList *ListNode
	for _,c := range val {
		tmpList := &ListNode{
			Val: int(c-'0'),
		}
		if retList == nil {
			retList = tmpList
		} else {
			tmpList.Next = retList
			retList = tmpList
		}
	}
	return retList
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val1, val2 := 0, 0
	for l1 != nil {
		val1 = val1*10 + l1.Val
		l1 = l1.Next
	}
	for l2 != nil {
		val2 = val2*10 + l2.Val
		l2 = l2.Next
	}
	result :=  val1 + val2
	return createList(strconv.Itoa(result))
}

func print(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Println()
}


func TestaddTwoNumbers() {
	l1 := createList("5")
	l2 := createList("5")
	retList := addTwoNumbers(l1, l2)
	print(retList)

	l1 = createList("243")
	l2 = createList("564")
	retList = addTwoNumbers(l1, l2)
	print(retList)

	l1 = createList("18")
	l2 = createList("0")
	retList = addTwoNumbers(l1, l2)
	print(retList)

	l1 = createList("98")
	l2 = createList("1")
	retList = addTwoNumbers(l1, l2)
	print(retList)

	l1 = createList("18")
	l2 = createList("0")
	retList = addTwoNumbers(l1, l2)
	print(retList)
}
