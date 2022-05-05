package title

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/

func indexChar(arr []rune, b rune) int {
	for index, item := range arr {
		if item == b {
			return index
		}
	}
	return -1
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	allChar := []rune{'(', '[', '{', ')', ']', '}'}
	allMap := []rune{'(': ')', '[': ']', '{': '}'}
	buf := []rune{}
	curIndex := -1
	for _, c := range s {
		index := indexChar(allChar, c)
		if index == -1 {
			return false
		}
		if index >= 3 {
			if curIndex >= 0 && allMap[buf[curIndex]] == c {
				curIndex--
			} else {
				return false
			}
		} else {
			curIndex++
			if curIndex >= len(buf) {
				buf = append(buf, c)
			} else {
				buf[curIndex] = c
			}
		}
	}
	return curIndex == -1
}

func TestIsValide() {
	tests := []struct {
		s    string
		isOk bool
	}{
		{
			"()[]{{}",
			false,
		},
		{
			"()[]",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"",
			true,
		},
		{
			"1",
			false,
		},
		{
			"()",
			true,
		},
		{
			"(",
			false,
		},
		{
			"(]",
			false,
		},
		{
			"(})[]{{}",
			false,
		},
		{
			"(o})[]{{}",
			false,
		},
		{
			"()",
			true,
		},
		{
			"[]",
			true,
		},
		{
			"{}",
			true,
		},
		{
			"[]{}",
			true,
		},
		{
			"(((((((((((((((())))))))))))))))",
			true,
		},
		{
			"[[[[[[[[[[[[]]]]]]]]]]]]",
			true,
		},
		{
			"{{{{{{{{{{{{}}}}}}}}}}}}",
			true,
		},
		{
			"{{{{{{{{{{{{}}}}}}}}}}}}[[[[[[[[[[[[]]]]]]]]]]]](((((((((((((((())))))))))))))))",
			true,
		},
		{
			"{{{{{{{{{{{{[[[[[[[[[[[[(((((((())))))))]]]]]]]]]]]]}}}}}}}}}}}}",
			true,
		},
		{
			"(((((((]())))))))",
			false,
		},
		{
			"(((((((())))))))[][][][][][][][][][]{}{}{}{}{}{}{}{}{}{}",
			true,
		},
	}
	for _, test := range tests {
		fmt.Println(isValid(test.s) == test.isOk)
	}
}
