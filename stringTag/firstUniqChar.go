package stringTag

import (
	"fmt"
)

/**
387. 字符串中的第一个唯一字符
 */

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	type dataItem struct {
		count int
		index int
	}

	buffer := []dataItem{'a': {}, 'z': {}}
	for index, c := range s {
		buffer[c].count++
		if buffer[c].count == 1 {
			buffer[c].index = index
		}
	}
	index := 1<<63 - 1
	for i := 'a'; i <= 'z'; i++ {
		if buffer[i].count == 1 && buffer[i].index < index {
			index = buffer[i].index
		}
	}
	if index != 1<<63-1 {
		return index
	}
	return -1
}

func TestfirstUniqChar() {
	tests := []struct {
		s      string
		result int
	}{
		{
			"",
			-1,
		},
		{
			"leetcode",
			0,
		},
		{
			"leetcodeleetcode",
			-1,
		},
		{
			"loveleetcodev",
			7,
		},
		{
			"loveleetcode",
			2,
		},
	}
	for _, test := range tests {
		fmt.Println(firstUniqChar(test.s) == test.result)
	}
}
