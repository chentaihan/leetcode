package mapTag

import "fmt"

/**
剑指 Offer II 017. 含有所有字符的最短字符串
给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。

如果 s 中存在多个符合条件的子字符串，返回任意一个。



注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。



示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最短子字符串 "BANC" 包含了字符串 t 的所有字符 'A'、'B'、'C'
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3：

输入：s = "a", t = "aa"
输出：""
解释：t 中两个字符 'a' 均应包含在 s 的子串中，因此没有符合条件的子字符串，返回空字符串。


提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成


进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/M1oyTv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	m2 := make([]int, 128)
	for i := 0; i < len(t); i++ {
		m2[t[i]]++
	}
	m1 := make([]int, 128)
	for i := 0; i < len(t); i++ {
		m1[s[i]]++
	}
	if arrayContain(m1, m2) {
		return s[:len(t)]
	}
	index := 0
	result := ""
	exist := false
	for i := len(t); i < len(s); i++ {
		m1[s[i]]++
		for arrayContain(m1, m2) {
			val := s[index : i+1]
			if !exist {
				result = val
				exist = true
			} else if len(val) < len(result) {
				result = val
			}
			m1[s[index]]--
			index++
		}
	}
	return result
}

func arrayContain(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] < arr2[i] {
			return false
		}
	}
	return true
}

func TestminWindow() {
	tests := []struct {
		s   string
		t   string
		res string
	}{
		{
			"ADOBECODEBANC", "ABC", "BANC",
		},
		{
			"ADOBECODEBAC", "ABC", "BAC",
		},
		{
			"ACOBECODEBAC", "ABC", "BAC",
		},
		{
			"ACOBECODEBaC", "ABC", "ACOB",
		},
		{
			"a", "a", "a",
		},
		{
			"abc", "cba", "abc",
		},
		{
			"a", "aa", "",
		},
	}
	for _, test := range tests {
		res := minWindow(test.s, test.t)
		if res != test.res {
			fmt.Println(res, test.res)
		} else {
			fmt.Println(res == test.res)
		}

	}
}
