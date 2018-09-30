package title

import "fmt"

/**
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
 */

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	l := len(s)
	curChar := s[0]
	for i := 1; i <= l/2; i++ {
		if l%i != 0 || curChar != s[i] {
			continue
		}
		j := 0
		//从0开始，验证所有下标相差i的元素值都一样
		for ; j < i; j++ {
			index := j + i
			for index < l {
				if s[j] == s[index] {
					index += i
				} else {
					goto tag
				}
			}
		}
		if j == i {
			return true
		}
	tag:
	}
	return false
}

func TestRepeatedSubstringPattern() {
	tests := []struct {
		s      string
		result bool
	}{
		{
			"",
			false,
		},
		{
			"a",
			false,
		},
		{
			"aa",
			true,
		},
		{
			"aaaaaaaaaaaaaaaaaa",
			true,
		},
		{
			"aaaaaaaaaaaaaaaaaaa",
			true,
		},
		{
			"aaaaaaaaaaaaaaaaaaabbbbbbbbbbaaaaaaaaaaaaaaaaaaabbbbbbbbbb",
			true,
		},
		{
			"aaaaaaaaaaaaaaaaaaabbbbbbbbbbaaaaaaaaaaaaaaaaaaabbbbbbbbbbb",
			false,
		},
		{
			"abab",
			true,
		},
		{
			"ababa",
			false,
		},
		{
			"abcabc",
			true,
		},
		{
			"abcabcd",
			false,
		},
		{
			"abccabcc",
			true,
		},
		{
			"abccabccs",
			false,
		},
		{
			"abbcabbc",
			true,
		},
		{
			"abbcdabbcde",
			false,
		},
		{
			"abbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcd",
			true,
		},
		{
			"abbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcdabbcd1",
			false,
		},
		{
			"abcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefg",
			true,
		},
		{
			"abcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgabcdefgh",
			false,
		},
	}
	for _, test := range tests {
		fmt.Println(repeatedSubstringPattern(test.s) == test.result)
	}
}
