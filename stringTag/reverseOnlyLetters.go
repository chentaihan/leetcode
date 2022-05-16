package stringTag

import "fmt"

/**
917. 仅仅反转字母
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。



示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"


提示：

S.length <= 100
33 <= S[i].ASCIIcode <= 122
S 中不包含 \ or "

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseOnlyLetters(S string) string {
	buf := make([]byte, len(S))
	start, end := 0, len(S)-1
	for start < end {
		for start < end && !isChar(S[start]) {
			buf[start] = S[start]
			start++
		}
		for start < end && !isChar(S[end]) {
			buf[end] = S[end]
			end--
		}
		if start < end {
			buf[start], buf[end] = S[end], S[start]
			start++
			end--
		}
	}
	if start == end {
		buf[start] = S[start]
	}
	return string(buf)
}

func isChar(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func TestreverseOnlyLetters() {
	tests := []struct {
		S string
		R string
	}{
		{
			"a",
			"a",
		},
		{
			"ab-cd",
			"dc-ba",
		},
		{
			"abcd",
			"dcba",
		},
		{
			"---",
			"---",
		},
		{
			"-ab-cd",
			"-dc-ba",
		},
		{
			"-ab-cd-",
			"-dc-ba-",
		},
	}

	for _, test := range tests {
		R := reverseOnlyLetters(test.S)
		if R != test.R {
			fmt.Println(R, test.R)
		}
	}
}
