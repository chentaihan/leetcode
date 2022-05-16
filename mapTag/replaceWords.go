package mapTag

import "strings"

/**
648. 单词替换
在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。

你需要输出替换之后的句子。

示例：

输入：dict(词典) = ["cat", "bat", "rat"] sentence(句子) = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

提示：

输入只包含小写字母。
1 <= dict.length <= 1000
1 <= dict[i].length <= 100
1 <= 句中词语数 <= 1000
1 <= 句中词语长度 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/replace-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func replaceWords(dict []string, sentence string) string {
	maxLen := 0
	m := map[string]struct{}{}
	for i := 0; i < len(dict); i++ {
		m[dict[i]] = struct{}{}
		if len(dict[i]) > maxLen {
			maxLen = len(dict[i])
		}
	}
	array := strings.Split(sentence, " ")
	for index, item := range array {
		min := maxLen
		if len(item) < min {
			min = len(item)
		}
		for i := 1; i <= min; i++ {
			if _, exist := m[item[:i]]; exist {
				array[index] = item[:i]
				break
			}
		}
	}
	return strings.Join(array, " ")
}
