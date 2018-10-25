package stringTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
49. 字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
分析：异位词排序后结果一样，map[string][]string
 */

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))
	for _, str := range strs {
		key := strSort(str)
		m[key] = append(m[key], str)
	}
	ret := make([][]string, 0, len(strs))
	for _, list := range m {
		ret = append(ret, list)
	}
	return ret
}

func groupAnagramsEx(strs []string) [][]string {
	m := make(map[string]int, len(strs))
	ret := make([][]string, 0, len(strs))
	index := 0
	for _, str := range strs {
		key := strSort(str)
		if i, exist := m[key]; exist {
			ret[i] = append(ret[i], str)
		} else {
			ret = append(ret, []string{str})
			m[key] = index
			index++
		}
	}
	return ret
}

func strSort(str string) string {
	buffer := [26]int{}
	for _, c := range str {
		buffer[c-'a']++
	}
	ret := make([]byte, len(str))
	index := 0
	for i, count := range buffer {
		if count > 0 {
			c := byte('a' + i)
			for j := 0; j < count; j++ {
				ret[index] = c
				index++
			}
		}
	}
	return string(ret)
}

func TestgroupAnagrams() {
	tests := []struct {
		strs []string
		ret  [][]string
	}{
		{
			[]string{"a", "a", "a", "ret", "ert", "tre", "a", "ewq"},
			[][]string{{"a", "a", "a", "a"}, {"ret", "ert", "tre"}, {"ewq"}},
		},
		{
			[]string{"abc", "bac", "cba", "ret", "ert", "tre", "qwe", "ewq"},
			[][]string{{"abc", "bac", "cba"}, {"ret", "ert", "tre"}, {"qwe", "ewq"}},
		},
	}
	for _, test := range tests {
		ret := groupAnagramsEx(test.strs)
		fmt.Println(common.StringArrayEqualEx(ret, test.ret))
	}
}
