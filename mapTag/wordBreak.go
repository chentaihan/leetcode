package mapTag

import (
	"fmt"
)

/**
139. 单词拆分
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	dict := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func TestWordBreak() {
	tests := []struct {
		s        string
		wordDict []string
		result   bool
	}{
		{
			"leetcode",
			[]string{"1", "leetcode"},
			true,
		},
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"leetcode1",
			[]string{"leet", "1", "code"},
			true,
		},
		{
			"qqqqqq",
			[]string{"q", "qq", "cc"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
		{
			"catsanddog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			true,
		},
	}
	for _, test := range tests {
		fmt.Println(wordBreak(test.s, test.wordDict) == test.result)
	}
}
