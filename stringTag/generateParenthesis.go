package stringTag

/**
22. 括号生成
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
https://leetcode-cn.com/problems/generate-parentheses/description/
*/

func GenerateParenthesis(n int) []string {
	var result []string
	if n <= 0 {
		return result
	}
	buffer := make([]byte, 0, n)
	generateParenthesis(buffer, &result, n, n)
	return result
}

func generateParenthesis(buffer []byte, result *[]string, left, right int) {
	if left == 0 && right == 0 {
		*result = append(*result, string(buffer))
		return
	}
	if left > 0 {
		generateParenthesis(append(buffer, '('), result, left-1, right)
	}
	if left < right && right > 0 {
		generateParenthesis(append(buffer, ')'), result, left, right-1)
	}
}
