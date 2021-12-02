package mapTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
140. 单词拆分 II
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：

分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
示例 2：

输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
示例 3：

输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func wordBreak1(s string, wordDict []string) []string {
	dict := make(map[string]bool, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	var result []string
	_wordBreak1(s, "", 0, dict, &result)
	return result
}

func _wordBreak1(s, list string, start int, dict map[string]bool, result *[]string) {
	for j := start + 1; j <= len(s); j++ {
		str := s[start:j]
		if dict[str] {
			isFirst := len(list) == 0
			if isFirst {
				list = str
			} else {
				list += " " + str
			}
			if j == len(s) {
				*result = append(*result, list)
				return
			} else {

				_wordBreak1(s, list, j, dict, result)
				if isFirst {
					list = ""
				} else {
					list = list[:len(list)-len(str)-1]
				}
			}
		}
	}
}

func TestwordBreak1() {
	tests := []struct {
		s      string
		dict   []string
		result []string
	}{
		{
			"a",
			[]string{"a"},
			[]string{"a"},
		},
		{
			"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cats and dog", "cat sand dog"},
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			[]string{},
		},
		{
			"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
	}

	for _, test := range tests {
		result := wordBreak1(test.s, test.dict)
		if common.StringEqualEx(result, test.result) {
			fmt.Println(true)
		} else {
			fmt.Println(false, result, test.result)
		}
	}
}
