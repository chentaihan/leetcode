package stringTag

import "fmt"

/**
1071. 字符串的最大公因子
对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。



示例 1：

输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
示例 2：

输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"
示例 3：

输入：str1 = "LEET", str2 = "CODE"
输出：""

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func TestgcdOfStrings() {
	tests := []struct {
		str1   string
		str2   string
		result string
	}{
		{
			"leet",
			"code",
			"",
		},
		{
			"",
			"",
			"",
		},
		{
			"AA",
			"A",
			"A",
		},
		{
			"AAAAAA",
			"AAA",
			"AAA",
		},
		{
			"AAAAAAAAA",
			"AAAAAA",
			"AAA",
		},
		{
			"ABAB",
			"AB",
			"AB",
		},
		{
			"ABCABC",
			"ABC",
			"ABC",
		},
		{
			"ABCDABCD",
			"ABCD",
			"ABCD",
		},
	}
	for _, test := range tests {
		result := gcdOfStrings(test.str1, test.str2)
		fmt.Println(result == test.result)
	}
}
