package dynamicTag

import "fmt"

/**
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' {
			stack = append(stack, i)
		} else if s[stack[len(stack)-1]] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, i)
		}
	}
	if len(stack) == 0 {
		return len(s)
	}
	maxLen := stack[0]
	for i := 1; i < len(stack); i++ {
		if stack[i]-stack[i-1]-1 > maxLen {
			maxLen = stack[i] - stack[i-1] - 1
		}
	}
	if len(s)-1-stack[len(stack)-1] > maxLen {
		maxLen = len(s) - 1 - stack[len(stack)-1]
	}
	return maxLen
}

func TestLongestValidParentheses() {
	tests := []struct {
		s   string
		res int
	}{
		{
			"(()", 2,
		},
		{
			"(())", 4,
		},
		{
			"()", 2,
		},
		{
			")(", 0,
		},
		{
			")()())", 4,
		},
		{
			"()())", 4,
		},
		{
			"()())(", 4,
		},
		{
			"(()())", 6,
		},
		{
			")))(((", 0,
		},
	}
	for _, test := range tests {
		res := longestValidParentheses(test.s)
		fmt.Println(res == test.res)
	}
}
