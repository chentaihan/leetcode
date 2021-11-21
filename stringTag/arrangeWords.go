package stringTag

import (
	"fmt"
	"sort"
)

/**
1451. 重新排列句子中的单词
「句子」是一个用空格分隔单词的字符串。给你一个满足下述格式的句子 text :

句子的首字母大写
text 中的每个单词都用单个空格分隔。
请你重新排列 text 中的单词，使所有单词按其长度的升序排列。如果两个单词的长度相同，则保留其在原句子中的相对顺序。

请同样按上述格式返回新的句子。

 

示例 1：

输入：text = "Leetcode is cool"
输出："Is cool leetcode"
解释：句子中共有 3 个单词，长度为 8 的 "Leetcode" ，长度为 2 的 "is" 以及长度为 4 的 "cool" 。
输出需要按单词的长度升序排列，新句子中的第一个单词首字母需要大写。
示例 2：

输入：text = "Keep calm and code on"
输出："On and keep calm code"
解释：输出的排序情况如下：
"On" 2 个字母。
"and" 3 个字母。
"keep" 4 个字母，因为存在长度相同的其他单词，所以它们之间需要保留在原句子中的相对顺序。
"calm" 4 个字母。
"code" 4 个字母。
示例 3：

输入：text = "To be or not to be"
输出："To be or to be not"
 

提示：

text 以大写字母开头，然后包含若干小写字母以及单词间的单个空格。
1 <= text.length <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rearrange-words-in-a-sentence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type item struct {
	str string
	pos int
}

func arrangeWords(text string) string {
	array := make([]item, 0, len(text)/100)
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			continue
		}
		for j := i; j < len(text); j++ {
			if text[j] == ' ' {
				array = append(array, item{text[i:j], len(array)})
				i = j
				break
			} else if j == len(text)-1 {
				array = append(array,item{ text[i:], len(array)})
				i = j
				break
			}
		}
	}
	sort.Sort(stringLen(array))
	result := make([]byte, len(text))
	copy(result, array[0].str)
	if array[0].str[0] >= 'a' && array[0].str[0] <= 'z' {
		result[0] = result[0] - 'a' + 'A'
	}
	l := len(array[0].str)
	for i := 1; i < len(array); i++ {
		result[l] = ' '
		l++
		copy(result[l:], array[i].str)
		if array[i].str[0] >= 'A' && array[i].str[0] <= 'Z' {
			result[l] = result[l] - 'A' + 'a'
		}
		l += len(array[i].str)
	}
	return string(result)
}

type stringLen []item

func (sl stringLen) Len() int {
	return len(sl)
}

func (sl stringLen) Less(i, j int) bool {
	if len(sl[i].str) == len(sl[j].str) {
		return sl[i].pos < sl[j].pos
	}
	return len(sl[i].str) < len(sl[j].str)
}

func (sl stringLen) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

func TestArrangeWords() {
	tests := []struct {
		text string
		res  string
	}{
		{
			"Jtik hrzrvhbmk gbo cfhmiqwonj ojkew ufvledv bomoeqt ops jgi zdhvbpbb zczmepdfpm jry rjazc titttcu",
			"Gbo ops jgi jry jtik ojkew rjazc ufvledv bomoeqt titttcu zdhvbpbb hrzrvhbmk cfhmiqwonj zczmepdfpm",
		},
		{
			"You and i",
			"I you and",
		},
		{
			"Leetcode is cool",
			"Is cool leetcode",
		},
		{
			"Keep calm and code on",
			"On and keep calm code",
		},
		{
			"To be or not to be",
			"To be or to be not",
		},
	}
	for _, test := range tests {
		res := arrangeWords(test.text)
		if res == test.res {
			fmt.Println(true)
		} else {
			fmt.Println(false, res, "   ", test.text)
		}
	}
}
