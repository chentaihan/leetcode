package stringTag

/**
242. 有效的字母异位词
分析：先统计第一个字符串中每个字符的数量，即每个字母++，遍历第二个字符串，每个字母--，最后每个字母的数量为0即为异位词
 */

func isAnagram(s string, t string) bool {
	buffer := [26]int{}
	for _, c := range s {
		buffer[c-'a']++
	}
	for _, c := range t {
		buffer[c-'a']--
	}
	for _, count := range buffer {
		if count != 0 {
			return false
		}
	}
	return true
}
