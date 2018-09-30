package title

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	if minLen == 0 {
		return ""
	}

	index := 0
	for ; index < minLen; index++ {
		curChar := strs[0][index]
		for _, str := range strs {
			if curChar != str[index] {
				goto ret
			}
		}
	}
ret:
	return strs[0][:index]
}

func TestLongestCommonPrefix() {
	fmt.Println(longestCommonPrefix(nil) == "")
	fmt.Println(longestCommonPrefix(nil) == "")
	fmt.Println(longestCommonPrefix([]string{""}) == "")
	fmt.Println(longestCommonPrefix([]string{"", ""}) == "")
	fmt.Println(longestCommonPrefix([]string{"a", "a"}) == "a")
	fmt.Println(longestCommonPrefix([]string{"a", "a", "a"}) == "a")
	fmt.Println(longestCommonPrefix([]string{"ab", "ab", "ab"}) == "ab")
	fmt.Println(longestCommonPrefix([]string{"ab", "abg", "ab"}) == "ab")
	fmt.Println(longestCommonPrefix([]string{"abg", "abg", "ab"}) == "ab")
	fmt.Println(longestCommonPrefix([]string{"a", "al"}) == "a")
	fmt.Println(longestCommonPrefix([]string{"al", "a"}) == "a")
	fmt.Println(longestCommonPrefix([]string{"abcdef", "abcdef"}) == "abcdef")
	fmt.Println(longestCommonPrefix([]string{"abcdef", "abcdefg"}) == "abcdef")
	fmt.Println(longestCommonPrefix([]string{"abcdefgh", "abcdefg"}) == "abcdefg")
	fmt.Println(longestCommonPrefix([]string{"sabcdefgh", "gabcdefg"}) == "")
}
