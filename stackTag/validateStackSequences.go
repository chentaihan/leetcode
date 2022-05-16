package stackTag

import "fmt"

/**
面试题31. 栈的压入、弹出序列
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。



示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func validateStackSequences(pushed []int, popped []int) bool {
	if len(popped) != len(popped) {
		return false
	}
	if len(popped) == 0 {
		return true
	}
	queue := []int{pushed[0]}
	pushIndex, popIndex := 1, 0
	for pushIndex < len(pushed) {
		if len(queue) == 0 {
			queue = append(queue, pushed[pushIndex])
			pushIndex++
		}
		if queue[len(queue)-1] == popped[popIndex] {
			queue = queue[:len(queue)-1]
			popIndex++
			continue
		}
		for pushIndex < len(pushed) && queue[len(queue)-1] != popped[popIndex] {
			queue = append(queue, pushed[pushIndex])
			pushIndex++
		}
	}
	queueIndex := len(queue) - 1
	for popIndex < len(popped) && queueIndex >= 0 {
		if popped[popIndex] != queue[queueIndex] {
			return false
		}
		queueIndex--
		popIndex++
	}
	return true
}

func TestvalidateStackSequences() {
	tests := []struct {
		push   []int
		pop    []int
		result bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 3, 1, 5, 4},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 1, 5, 4, 3},
			true,
		},
	}
	for index, test := range tests {
		result := validateStackSequences(test.push, test.pop)
		if result != test.result {
			fmt.Println("error", index)
		}
	}
}
