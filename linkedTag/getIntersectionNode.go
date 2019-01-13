package linkedTag

import "github.com/chentaihan/leetcode/common"

/**
160. 相交链表
 */

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	countA := common.GetListNodeCount(headA)
	countB := common.GetListNodeCount(headB)
	if countA == 0 || countB == 0 {
		return nil
	}
	if countA > countB {
		for countA > countB {
			headA = headA.Next
			countA--
		}
	} else {
		for countB > countA {
			headB = headB.Next
			countB--
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
