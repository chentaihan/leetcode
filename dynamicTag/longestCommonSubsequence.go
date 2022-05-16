package dynamicTag

import "fmt"

/**
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。



示例 1:

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。


提示:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：非常经典的动态规划问题，对于给定字符串text1的位置iii和字符串text2的位置jjj，定义f(i,j)f(i,j)f(i,j)表示text[:i-1]和text[:j-1]的最长公共子序列。如果此时text1[i-1] == text2[j-1]，那么
f(i,j)=f(i−1,j−1)+1
如果不相等，那么此时的最大值应该来自两部分：text[:i]和text[:j-1]、text[:i-1]和text[:j]。
f(i,j)=max(f(i,j−1),f(i−1,j))
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	row, col := len(text1), len(text2)
	buf := make([][]int, 0, row)
	for i := 0; i < row; i++ {
		buf = append(buf, make([]int, col))
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			val := 0
			if text1[i] == text2[j] {
				if i > 0 && j > 0 {
					val = buf[i-1][j-1] + 1
				} else {
					val = 1
				}
			} else {
				if i > 0 {
					val = buf[i-1][j]
				}
				if j > 0 && buf[i][j-1] > val {
					val = buf[i][j-1]
				}
			}

			buf[i][j] = val
		}
	}
	return buf[row-1][col-1]
}

func TestlongestCommonSubsequence() {
	tests := []struct {
		text1 string
		text2 string
		count int
	}{
		{
			"aaaaaaaaaa",
			"aaaaaaa",
			7,
		},
		{
			"aaaaaaaaaa",
			"aaabaaaa",
			7,
		},
		{
			"aaaaaaaaaa",
			"aaabaacaa",
			7,
		},
		{
			"aaaaaaaaaa",
			"aafabaacaa",
			7,
		},
		{
			"aaaaaaaaaa",
			"asafabaacaa",
			7,
		},
		{
			"abcdefg",
			"aceg",
			4,
		},
		{
			"abcdefgjkl",
			"aecegjkl",
			7,
		},
		{
			"a",
			"a",
			1,
		},
		{
			"ab",
			"ab",
			2,
		},
		{
			"abc",
			"abc",
			3,
		},
		{
			"abcd",
			"abcd",
			4,
		},
		{
			"abc",
			"abcdefghjkl",
			3,
		},
		{
			"abcdefghjkl",
			"abc",
			3,
		},
		{
			"abcdefghjkl",
			"abc",
			3,
		},
	}
	for index, test := range tests {
		count := longestCommonSubsequence(test.text1, test.text2)
		fmt.Println(index, count == test.count)
	}
}
