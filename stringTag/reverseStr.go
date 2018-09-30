package stringTag

import "fmt"

/**
541. 反转字符串 II
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。
如果剩余少于 k 个字符，则将剩余的所有全部反转。
如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。
 */

func reverseStr(s string, k int) string {
	if k <= 1 {
		return s
	}
	l := len(s)
	buffer := make([]byte, l)
	for i := 0; i < l; i++ {
		buffer[i] = s[i]
	}
	status := false
	for i := 0; i < l; i += k {
		if !status {
			end := i+k
			if end > l {
				end = l
			}
			reverseK(buffer, i, end)
		}
		status = !status
	}
	return string(buffer)
}

func reverseK(buf []byte, start, end int) {
	for end--; start < end; {
		buf[start], buf[end] = buf[end], buf[start]
		start++
		end--
	}
}

func TestreverseStr() {
	tests := []struct {
		s      string
		k      int
		result string
	}{
		{
			"abcdefg",
			8,
			"gfedcba",
		},
		{
			"abcdefghij",
			10,
			"jihgfedcba",
		},
		{
			"abcdefg",
			1,
			"abcdefg",
		},
		{
			"abcdefg",
			2,
			"bacdfeg",
		},
		{
			"abcdefg",
			3,
			"cbadefg",
		},
		{
			"abcdefghi",
			3,
			"cbadefihg",
		},
		{
			"abcdefghik",
			3,
			"cbadefihgk",
		},
		{
			"abcdefghikl",
			3,
			"cbadefihgkl",
		},
		{
			"abcdefghiklm",
			3,
			"cbadefihgklm",
		},
		{
			"abcdefghiklmn",
			3,
			"cbadefihgklmn",
		},
		{
			"abcdefghiklmno",
			3,
			"cbadefihgklmon",
		},
		{
			"abcdefghiklmnop",
			3,
			"cbadefihgklmpon",
		},
	}
	for _, test := range tests {
		fmt.Println(reverseStr(test.s, test.k) == test.result)
	}
}
