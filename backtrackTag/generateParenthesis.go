package backtrackTag

import "fmt"

/**
面试题 08.09. 括号
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bracket-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateParenthesis(n int) []string {
	result := make([]string, 0, n)
	list := make([]byte, 0, n*2)
	_generateParenthesis(n, n, &result, list)
	return result
}

func _generateParenthesis(left, right int, result *[]string, list []byte) {
	if left == 0 && right == 0 {
		*result = append(*result, string(list))
		return
	}
	if left > 0 {
		list = append(list, '(')
		_generateParenthesis(left-1, right, result, list)
		list = list[:len(list)-1]
	}
	if left < right {
		list = append(list, ')')
		_generateParenthesis(left, right-1, result, list)
		list = list[:len(list)-1]
	}
}

func TestgenerateParenthesis() {
	array := generateParenthesis(3)
	fmt.Println(array)
}
