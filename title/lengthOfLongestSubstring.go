package title

import "fmt"

/**
剑指 Offer II 016. 不含重复字符的最长子字符串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0


提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wtcaE1
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[rune]int{}
	maxLen, start := 0, 0
	for index, c := range s {
		existValue, exist := m[c]
		if index-start > maxLen {
			if exist && existValue >= start {
				maxLen = index - start - 1
			} else {
				maxLen = index - start
			}
		}
		if exist {
			if existValue >= start {
				start = existValue + 1
			}
		}
		m[c] = index
	}
	return maxLen + 1
}

func TestLengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("") == 0)
	fmt.Println(lengthOfLongestSubstring("b") == 1)
	fmt.Println(lengthOfLongestSubstring("ab") == 2)
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("dvdf") == 3)
	fmt.Println(lengthOfLongestSubstring("tmmzuxt") == 5)
}
