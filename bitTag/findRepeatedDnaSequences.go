package bitTag

/**
187. 重复的DNA序列
所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

编写一个函数来查找 DNA 分子中所有出现超多一次的10个字母长的序列（子串）。

示例:

输入: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

输出: ["AAAAACCCCC", "CCCCCAAAAA"]
 */

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	m := map[string]int{}
	for i := 0; i <= len(s)-10; i++ {
		m[s[i:i+10]]++
	}
	var ret []string
	for key, count := range m {
		if count > 1 {
			ret = append(ret, key)
		}
	}
	return ret
}
