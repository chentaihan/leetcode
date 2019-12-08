package stringTag

/**
58. 最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:

输入: "Hello World"
输出: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func lengthOfLastWord(s string) int {
	end := -1
	for index := len(s) - 1; index >= 0; index-- {
		if end == -1 && s[index] == ' ' {
			continue
		}
		if end == -1 {
			end = index
		} else if s[index] == ' ' {
			return end - index
		}
	}
	if end > -1 {
		return end + 1
	}
	return 0
}
