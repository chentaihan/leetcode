package stringTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
821. 字符的最短距离
给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

示例 1:

输入: S = "loveleetcode", C = 'e'
输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
*/

func shortestToChar(S string, C byte) []int {
	ret := make([]int, len(S))
	target := []int{}
	for i := 0; i < len(S); i++ {
		if C == S[i] {
			target = append(target, i)
		}
	}

	targetIndex := 0
	for i := 0; i < len(S); i++ {
		start := target[targetIndex]
		if i <= start {
			ret[i] = start - i
			continue
		}

		end := start
		if targetIndex+1 < len(target) {
			end = target[targetIndex+1]
		}

		if i < end {
			if i-start > end-i {
				ret[i] = end - i
			} else {
				ret[i] = i - start
			}
		} else {
			ret[i] = i - end
			if targetIndex+1 < len(target) {
				targetIndex++
			}
		}
	}
	return ret
}

func TestshortestToChar() {
	tests := []struct {
		S   string
		C   byte
		ret []int
	}{
		{
			"loveleetcode",
			'e',
			[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			"aaba",
			'b',
			[]int{2, 1, 0, 1},
		},
	}

	for _, test := range tests {
		ret := shortestToChar(test.S, test.C)
		fmt.Println(common.IntEqual(ret, test.ret))
	}
}
