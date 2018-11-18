package arrayTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
830. 较大分组的位置
在一个由小写字母构成的字符串 S 中，包含由一些连续的相同字符所构成的分组。

例如，在字符串 S = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。

我们称所有包含大于或等于三个连续字符的分组为较大分组。找到每一个较大分组的起始和终止位置。

最终结果按照字典顺序输出。

示例 1:

输入: "abbxxxxzzy"
输出: [[3,6]]
解释: "xxxx" 是一个起始于 3 且终止于 6 的较大分组。
示例 2:

输入: "abc"
输出: []
解释: "a","b" 和 "c" 均不是符合要求的较大分组。
示例 3:

输入: "abcdddeeeeaabbbcd"
输出: [[3,5],[6,9],[12,14]]
 */

func largeGroupPositions(S string) [][]int {
	var val uint8
	start, end := 0, 0
	result := [][]int{}
	for i := 0; i < len(S); i++ {
		if val != S[i] {
			if end-start >= 2 {
				result = append(result, []int{start, end})
			}
			start, end = i, i
			val = S[i]
		} else {
			end++
		}
	}
	if end-start >= 2 {
		result = append(result, []int{start, end})
	}
	return result
}

func TestlargeGroupPositions() {
	tests := []struct {
		s   string
		ret [][]int
	}{
		{
			"bb",
			[][]int{},
		},
		{
			"bbb",
			[][]int{{0, 2}},
		},
		{
			"bbbb",
			[][]int{{0, 3}},
		},
		{
			"aaaabbbb",
			[][]int{{0, 3},{4,7}},
		},
		{
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			[][]int{{0, 39}},
		},
		{
			"aaabbbcdddefffffgg",
			[][]int{{0, 2}, {3, 5}, {7, 9}, {11, 15}},
		},
		{
			"aaabbbcdddefffffggg",
			[][]int{{0, 2}, {3, 5}, {7, 9}, {11, 15}, {16, 18}},
		},
	}

	for _, test := range tests {
		ret := largeGroupPositions(test.s)
		fmt.Println(common.IntEqualEx(ret,test.ret))
	}
}
