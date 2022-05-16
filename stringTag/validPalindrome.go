package stringTag

import "fmt"

/**
680. 验证回文字符串 Ⅱ
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func validPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}
	return _validPalindrome(s, 1)
}

func _validPalindrome(s string, count int) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
			continue
		}
		if end-start <= 1 {
			return true
		}
		if count <= 0 {
			return false
		}
		a, b := s[start+1] == s[end], s[start] == s[end-1]
		if a || b {
			if a && !b {
				return _validPalindrome(s[start+1:end+1], count-1)
			} else if !a && b {
				return _validPalindrome(s[start:end], count-1)
			}
			return _validPalindrome(s[start+1:end+1], count-1) || _validPalindrome(s[start:end], count-1)
		} else {
			return false
		}
	}
	return true
}

func TestvalidPalindrome() {
	tests := []struct {
		s      string
		result bool
	}{
		{
			"eeccccbebaeeabebccceea",
			false,
		},
		{
			"a121",
			true,
		},
		{
			"121a",
			true,
		},
		{
			"12",
			true,
		},
		{
			"121",
			true,
		},
		{
			"1211",
			true,
		},
		{
			"123231",
			true,
		},
		{
			"aa121aa",
			true,
		},
		{
			"1234567890987654321",
			true,
		},
	}

	for _, test := range tests {
		result := validPalindrome(test.s)
		fmt.Println(result == test.result)
	}
}
