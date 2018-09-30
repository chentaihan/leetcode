package title

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[rune]int{}
	maxLen, start := 0, 0
	for index, c := range s {
		existValue, exist := m[c]
		if index-start > maxLen {
			if exist && existValue >= start {
				maxLen = index - start - 1
			} else {
				maxLen = index - start
			}
		}
		if exist {
			if existValue >= start {
				start = existValue + 1
			}
		}
		m[c] = index
	}
	return maxLen + 1
}

func TestLengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("") == 0)
	fmt.Println(lengthOfLongestSubstring("b") == 1)
	fmt.Println(lengthOfLongestSubstring("ab") == 2)
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("dvdf") == 3)
	fmt.Println(lengthOfLongestSubstring("tmmzuxt") == 5)
}
