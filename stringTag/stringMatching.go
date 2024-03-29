package stringTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"strings"
)

/**
1408. 数组中的字符串匹配
给你一个字符串数组 words ，数组中的每个字符串都可以看作是一个单词。请你按 任意 顺序返回 words 中是其他单词的子字符串的所有单词。

如果你可以删除 words[j] 最左侧和/或最右侧的若干字符得到 words[i] ，那么字符串 words[i] 就是 words[j] 的一个子字符串。



示例 1：

输入：words = ["mass","as","hero","superhero"]
输出：["as","hero"]
解释："as" 是 "mass" 的子字符串，"hero" 是 "superhero" 的子字符串。
["hero","as"] 也是有效的答案。
示例 2：

输入：words = ["leetcode","et","code"]
输出：["et","code"]
解释："et" 和 "code" 都是 "leetcode" 的子字符串。
示例 3：

输入：words = ["blue","green","bu"]
输出：[]


提示：

1 <= words.length <= 100
1 <= words[i].length <= 30
words[i] 仅包含小写英文字母。
题目数据 保证 每个 words[i] 都是独一无二的。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/string-matching-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func stringMatching(words []string) []string {
	var res []string
	m := make(map[string]struct{})
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if len(words[i]) < len(words[j]) {
				if strings.Index(words[j], words[i]) >= 0 {
					if _, exist := m[words[i]]; !exist {
						res = append(res, words[i])
						m[words[i]] = struct{}{}
					}
				}
			} else {
				if strings.Index(words[i], words[j]) >= 0 {
					if _, exist := m[words[j]]; !exist {
						res = append(res, words[j])
						m[words[j]] = struct{}{}
					}
				}
			}
		}
	}
	return res
}

func TestStringMatching() {
	tests := []struct {
		words []string
		res   []string
	}{
		{
			[]string{"mass", "as", "hero", "superhero"},
			[]string{"as", "hero"},
		},
		{
			[]string{"leetcode", "et", "code"},
			[]string{"et", "code"},
		},
		{
			[]string{"et", "code", "leetcode", "leetcode1"},
			[]string{"et", "code", "leetcode"},
		},
	}
	for _, test := range tests {
		res := stringMatching(test.words)
		if common.StringEqual(res, test.res) {
			fmt.Println(true)
		} else {
			fmt.Println(res, test.res)
		}
	}
}
