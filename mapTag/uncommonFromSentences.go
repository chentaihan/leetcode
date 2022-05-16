package mapTag

import "strings"

/**
884. 两句话中的不常见单词
给定两个句子 A 和 B 。 （句子是一串由空格分隔的单词。每个单词仅由小写字母组成。）

如果一个单词在其中一个句子中只出现一次，在另一个句子中却没有出现，那么这个单词就是不常见的。

返回所有不常用单词的列表。

您可以按任何顺序返回列表。



示例 1：

输入：A = "this apple is sweet", B = "this apple is sour"
输出：["sweet","sour"]
示例 2：

输入：A = "apple apple", B = "banana"
输出：["banana"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/uncommon-words-from-two-sentences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func uncommonFromSentences(A string, B string) []string {
	arrA := strings.Split(A, " ")
	arrB := strings.Split(B, " ")
	mA := map[string]int{}
	mB := map[string]int{}
	for _, item := range arrA {
		mA[item]++
	}
	for _, item := range arrB {
		mB[item]++
	}

	var result []string
	for key, count := range mA {
		if count == 1 && mB[key] == 0 {
			result = append(result, key)
		}
	}
	for key, count := range mB {
		if count == 1 && mA[key] == 0 {
			result = append(result, key)
		}
	}

	return result
}
