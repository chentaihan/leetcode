package mapTag

import "fmt"

/**
748. 最短完整词
如果单词列表（words）中的一个单词包含牌照（licensePlate）中所有的字母，那么我们称之为完整词。在所有完整词中，最短的单词我们称之为最短完整词。

单词在匹配牌照中的字母时不区分大小写，比如牌照中的 "P" 依然可以匹配单词中的 "p" 字母。

我们保证一定存在一个最短完整词。当有多个单词都符合最短完整词的匹配条件时取单词列表中最靠前的一个。

牌照中可能包含多个相同的字符，比如说：对于牌照 "PP"，单词 "pair" 无法匹配，但是 "supper" 可以匹配。

 

示例 1：

输入：licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
输出："steps"
说明：最短完整词应该包括 "s"、"p"、"s" 以及 "t"。对于 "step" 它只包含一个 "s" 所以它不符合条件。同时在匹配过程中我们忽略牌照中的大小写。
 

示例 2：

输入：licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
输出："pest"
说明：存在 3 个包含字母 "s" 且有着最短长度的完整词，但我们返回最先出现的完整词。
 

注意:

牌照（licensePlate）的长度在区域[1, 7]中。
牌照（licensePlate）将会包含数字、空格、或者字母（大写和小写）。
单词列表（words）长度在区间 [10, 1000] 中。
每一个单词 words[i] 都是小写，并且长度在区间 [1, 15] 中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-completing-word
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func shortestCompletingWord(licensePlate string, words []string) string {
	m := getshortestCompletingWord(licensePlate)
	result := ""
	for _, item := range words {
		mItem := getshortestCompletingWord(item)
		if len(m) > len(mItem) {
			continue
		}
		isOk := true
		for key, value := range m {
			if mItem[key] < value {
				isOk = false
				break
			}
		}
		if isOk {
			if result == "" {
				result = item
			} else if len(item) < len(result) {
				result = item
			}
		}
	}

	return result
}

func getshortestCompletingWord(words string) map[byte]int {
	m := map[byte]int{}
	for i := 0; i < len(words); i++ {
		if 'a' <= words[i] && words[i] <= 'z' {
			m[words[i]]++
		} else if 'A' <= words[i] && words[i] <= 'Z' {
			m[words[i]+32]++
		}
	}
	return m
}

func TestshortestCompletingWord() {
	tests := []struct {
		licensePlate string
		words        []string
		result       string
	}{
		{
			"1s3 PSt",
			[]string{"step", "steps", "stripe", "stepple"},
			"steps",
		},
		{
			"1s3 456",
			[]string{"looks", "pest", "stew", "show"},
			"pest",
		},
	}
	for _, test := range tests {
		result := shortestCompletingWord(test.licensePlate, test.words)
		fmt.Println(result == test.result)
	}
}
