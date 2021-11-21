package backtrackTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
131. 分割回文串
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-partitioning
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func partition(s string) [][]string {
	result := make([][]string, 0, len(s))
	list := make([]string, 0, len(s))
	_partition(s, 0, &result, list)
	return result
}

func _partition(s string, start int, result *[][]string, list []string) {
	if len(s) == start {
		newList := make([]string, len(list))
		copy(newList, list)
		*result = append(*result, newList)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			list = append(list, s[start:i+1])
			_partition(s, i+1, result, list)
			list = list[:len(list)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func TestPartition() {
	tests := []struct {
		text string
		res  [][]string
	}{
		{
			"aab",
			[][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			"a",
			[][]string{{"a"}},
		},
		{
			"aa",
			[][]string{{"a", "a"}, {"aa"}},
		},
	}
	for _, test := range tests {
		res := partition(test.text)
		fmt.Println(common.StringArrayEqualEx(res, test.res))
	}
}
