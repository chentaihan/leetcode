package stringTag

import "fmt"

/**
409. 最长回文串
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	sum, left, one := 0, 0, 0
	for _, count := range m {
		if count > 1 {
			sum += count
			if count%2 == 1 {
				left++
			}
		} else {
			one = 1
		}
	}
	if one < 1 && left >= 1 {
		left--
	}
	return sum - left + one
}

func TestlongestPalindrome() {
	tests := []struct {
		s     string
		count int
	}{
		{
			"aabbcccdd",
			9,
		},
		{
			"",
			0,
		},
		{
			"0",
			1,
		},
		{
			"01",
			1,
		},
		{
			"11",
			2,
		},
		{
			"aabbcccddeee",
			11,
		},
		{
			"aabbcccddeeef",
			11,
		},
		{
			"abccccdd",
			7,
		},
		{
			"aabbccdd",
			8,
		},
		{
			"aabbcccddd",
			9,
		},
		{
			"aabbcccddde",
			9,
		},
	}
	for _, test := range tests {
		count := longestPalindrome(test.s)
		fmt.Println(count == test.count)
	}
}
