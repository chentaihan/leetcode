package mapTag

/**
剑指 Offer II 014. 字符串中的变位词
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。

换句话说，第一个字符串的排列之一是第二个字符串的 子串 。



示例 1：

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").
示例 2：

输入: s1= "ab" s2 = "eidboaoo"
输出: False


提示：

1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/MPnaiL
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m1 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a']++
	}
	m2 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		m2[s2[i]-'a']++
	}
	for i := len(s1); i < len(s2); i++ {
		if arrayEqual(m1, m2) {
			return true
		}
		m2[s2[i-len(s1)]-'a']--
		m2[s2[i]-'a']++
	}
	return arrayEqual(m1, m2)
}

func arrayEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
