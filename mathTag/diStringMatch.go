package mathTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
942. 增减字符串匹配
给定只含 "I"（增大）或 "D"（减小）的字符串 S ，令 N = S.length。

返回 [0, 1, ..., N] 的任意排列 A 使得对于所有 i = 0, ..., N-1，都有：

如果 S[i] == "I"，那么 A[i] < A[i+1]
如果 S[i] == "D"，那么 A[i] > A[i+1]

示例 1：

输出："IDID"
输出：[0,4,1,3,2]
*/

func diStringMatch(S string) []int {
	start, end := 0, len(S)
	ret := make([]int, end+1)
	for index, s := range S {
		if s == 'I' {
			ret[index] = start
			start++
		} else {
			ret[index] = end
			end--
		}
	}
	ret[len(S)] = end
	return ret
}

func TestdiStringMatch() {
	tests := []struct {
		S   string
		ret []int
	}{
		{
			"IDID",
			[]int{0, 4, 1, 3, 2},
		},
		{
			"III",
			[]int{0, 1, 2, 3},
		},
		{
			"DDD",
			[]int{3, 2, 1, 0},
		},
		{
			"DDI",
			[]int{3, 2, 0, 1},
		},
		{
			"IDIDID",
			[]int{0, 6, 1, 5, 2, 4, 3},
		},
	}

	for _, test := range tests {
		ret := diStringMatch(test.S)
		fmt.Println(common.IntEqual(ret, test.ret))
	}
}
