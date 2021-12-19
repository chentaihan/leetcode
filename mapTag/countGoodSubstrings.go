package mapTag

/**
1876. 长度为三且各字符不同的子字符串
如果一个字符串不含有任何重复字符，我们称这个字符串为 好 字符串。

给你一个字符串 s ，请你返回 s 中长度为 3 的 好子字符串 的数量。

注意，如果相同的好子字符串出现多次，每一次都应该被记入答案之中。

子字符串 是一个字符串中连续的字符序列。



示例 1：

输入：s = "xyzzaz"
输出：1
解释：总共有 4 个长度为 3 的子字符串："xyz"，"yzz"，"zza" 和 "zaz" 。
唯一的长度为 3 的好子字符串是 "xyz" 。
示例 2：

输入：s = "aababcabc"
输出：4
解释：总共有 7 个长度为 3 的子字符串："aab"，"aba"，"bab"，"abc"，"bca"，"cab" 和 "abc" 。
好子字符串包括 "abc"，"bca"，"cab" 和 "abc" 。


提示：

1 <= s.length <= 100
s​​​​​​ 只包含小写英文字母。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/substrings-of-size-three-with-distinct-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	l := len(s) - 1
	count := 0
	for i := 1; i < l; i++ {
		if s[i-1] != s[i] && s[i] != s[i+1] && s[i-1] != s[i+1] {
			count++
		}
	}
	return count
}
