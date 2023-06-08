package mapTag

import "fmt"

/**
1935. 可以输入的最大单词数
键盘出现了一些故障，有些字母键无法正常工作。而键盘上所有其他键都能够正常工作。

给你一个由若干单词组成的字符串 text ，单词间由单个空格组成（不含前导和尾随空格）；另有一个字符串 brokenLetters ，由所有已损坏的不同字母键组成，返回你可以使用此键盘完全输入的 text 中单词的数目。



示例 1：

输入：text = "hello world", brokenLetters = "ad"
输出：1
解释：无法输入 "world" ，因为字母键 'd' 已损坏。
示例 2：

输入：text = "leet code", brokenLetters = "lt"
输出：1
解释：无法输入 "leet" ，因为字母键 'l' 和 't' 已损坏。
示例 3：

输入：text = "leet code", brokenLetters = "e"
输出：0
解释：无法输入任何单词，因为字母键 'e' 已损坏。


提示：

1 <= text.length <= 104
0 <= brokenLetters.length <= 26
text 由若干用单个空格分隔的单词组成，且不含任何前导和尾随空格
每个单词仅由小写英文字母组成
brokenLetters 由 互不相同 的小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-number-of-words-you-can-type
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func canBeTypedWords(text string, brokenLetters string) int {
	count := 0
	var m [26]bool
	for i := 0; i < len(brokenLetters); i++ {
		m[brokenLetters[i]-'a'] = true
	}
	var buf []byte
	for i := 0; i < len(text); i++ {
		if text[i] >= 'a' && text[i] <= 'z' {
			buf = append(buf, text[i])
		} else if len(buf) > 0 {
			j := 0
			for ; j < len(buf); j++ {
				if m[buf[j]-'a'] {
					break
				}
			}
			if j == len(buf) {
				count++
			}
			buf = buf[:0]
		}
	}
	j := 0
	for ; j < len(buf); j++ {
		if m[buf[j]-'a'] {
			break
		}
	}
	if j == len(buf) {
		count++
	}
	return count
}

func TestCanBeTypedWords() {
	tests := []struct {
		text          string
		brokenLetters string
		res           int
	}{
		{
			"leet code",
			"e",
			0,
		},
		{
			"hello world",
			"ad",
			1,
		},
		{
			"hello world",
			"f",
			2,
		},
		{
			"hello world",
			"hd",
			0,
		},
		{
			"leet code",
			"lt",
			1,
		},
	}
	for _, test := range tests {
		res := canBeTypedWords(test.text, test.brokenLetters)
		if res == test.res {
			fmt.Println(true)
		} else {
			fmt.Println(res, test.res)
		}
	}
}
