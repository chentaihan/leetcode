package stringTag

import "fmt"

/**
125. 验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if !_isPalindrome(s[start]) {
			start++
			continue
		}
		if !_isPalindrome(s[end]) {
			end--
			continue
		}

		c := s[start]
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		d := s[end]
		if d >= 'A' && d <= 'Z' {
			d += 32
		}
		if c == d {
			start++
			end--
			continue
		}
		return false
	}
	return true
}

func _isPalindrome(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func TestisPalindrome() {
	tests := []struct {
		s      string
		result bool
	}{
		{
			"0P",
			false,
		},
		{
			"OP",
			false,
		},
		{
			"OPO",
			true,
		},
		{
			"oOPOO",
			true,
		},
		{
			"1oOPOO 1",
			true,
		},
		{
			"1,oO,POO ,1",
			true,
		},
		{
			"A man, a plan, a canal: Panama",
			true,
		},
	}
	for _, test := range tests {
		result := isPalindrome(test.s)
		fmt.Println(result == test.result)
	}
}
