package stringTag

import "fmt"

/**
856. 括号的分数
给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

() 得 1 分。
AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
(A) 得 2 * A 分，其中 A 是平衡括号字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/score-of-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func scoreOfParentheses(S string) int {
	sum := 0
	var count uint
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			count++
		} else {
			count--
			sum += 1 << count
		}
	}
	return sum
}

func TestscoreOfParentheses() {
	tests := []struct {
		S string
		R int
	}{
		{
			"()",
			1,
		},
		{
			"()()",
			2,
		},
		{
			"(())",
			2,
		},
	}

	for _, test := range tests {
		R := scoreOfParentheses(test.S)
		fmt.Println(R == test.R)
	}
}
