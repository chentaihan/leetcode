package backtrackTag

/**
剑指 Offer 38. 字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]


限制：

1 <= s 的长度 <= 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permutation(s string) []string {
	result := make([]string, 0, len(s))
	_permutation(0, &result, []byte(s))
	return result
}

func _permutation(index int, result *[]string, list []byte) {
	if index == len(list) {
		*result = append(*result, string(list))
		return
	}
	m := make(map[byte]bool, len(list))
	for i := index; i < len(list); i++ {
		if !m[list[i]] {
			m[list[i]] = true
			list[i], list[index] = list[index], list[i]
			_permutation(index+1, result, list)
			list[i], list[index] = list[index], list[i]
		}
	}
}
