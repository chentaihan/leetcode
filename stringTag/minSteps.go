package stringTag

/**
1347. 制造字母异位词的最小步骤数 TODO
给你两个长度相等的字符串 s 和 t。每一个步骤中，你可以选择将 t 中的 任一字符 替换为 另一个字符。

返回使 t 成为 s 的字母异位词的最小步骤数。

字母异位词 指字母相同，但排列不同（也可能相同）的字符串。



示例 1：

输出：s = "bab", t = "aba"
输出：1
提示：用 'b' 替换 t 中的第一个 'a'，t = "bba" 是 s 的一个字母异位词。
示例 2：

输出：s = "leetcode", t = "practice"
输出：5
提示：用合适的字符替换 t 中的 'p', 'r', 'a', 'i' 和 'c'，使 t 变成 s 的字母异位词。
示例 3：

输出：s = "anagram", t = "mangaar"
输出：0
提示："anagram" 和 "mangaar" 本身就是一组字母异位词。
示例 4：

输出：s = "xxyyzz", t = "xxyyzz"
输出：0
示例 5：

输出：s = "friend", t = "family"
输出：4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minSteps(s string, t string) int {
	ms := make(map[byte]int)
	mt := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	//for key, count := range mt {
	//
	//}
	return 0
}
