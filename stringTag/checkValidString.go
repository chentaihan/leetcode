package stringTag

import "fmt"

/**
678. 有效的括号字符串
给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

任何左括号 ( 必须有相应的右括号 )。
任何右括号 ) 必须有相应的左括号 ( 。
左括号 ( 必须在对应的右括号之前 )。
* 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
一个空字符串也被视为有效字符串。
示例 1:

输入: "()"
输出: True
示例 2:

输入: "(*)"
输出: True
示例 3:

输入: "(*))"
输出: True
注意:

字符串大小将在 [1，100] 范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parenthesis-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：主要还是考虑顺序的问题，所以使用了一个基于数组的栈
*/

func checkValidString(s string) bool {
	cList := make([]int, 0, len(s)/2)
	xList := make([]int, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if len(cList) == 0 {
				if len(xList) == 0 {
					return false
				}
				xList = xList[:len(xList)-1]
			} else {
				cList = cList[:len(cList)-1]
			}
		} else if s[i] == '(' {
			cList = append(cList, i)
		} else {
			xList = append(xList, i)
		}
	}
	if len(cList) > 0 {
		if len(cList) > len(xList) {
			return false
		}
		index := len(xList) - len(cList)
		for i := 0; i < len(cList); i++ {
			if cList[i] > xList[index+i] {
				return false
			}
		}
	}

	return true
}

func TestcheckValidString() {
	tests := []struct {
		str    string
		result bool
	}{
		{
			"*****(((((",
			false,
		},
		{
			"((((()))))",
			true,
		},
		{
			"((()))(())()",
			true,
		},
		{
			"(((()))(())())",
			true,
		},
		{
			"((((()))(())())*",
			true,
		},
		{
			"(((((()))(())())*)",
			true,
		},
		{
			"(((((()))(())())*)*",
			true,
		},
		{
			"*(((((()))(())())*)*",
			true,
		},
		{
			"(())((())()()(*)(*()(())())())()()((()())((()))(*",
			false,
		},
	}
	for _, test := range tests {
		result := checkValidString(test.str)
		fmt.Println(result == test.result)
	}
}
