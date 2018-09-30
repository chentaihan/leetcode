package stringTag

import (
	"fmt"
)

/**
28. 实现strStr()
 */

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		nIndex := 0
		hIndex := i
		for nIndex < len(needle) && haystack[hIndex] == needle[nIndex] {
			hIndex++
			nIndex++
		}
		if nIndex == len(needle) {
			return i
		}
	}
	return -1
}

func TeststrStr() {
	tests := []struct {
		haystack string
		needle   string
		index    int
	}{
		{
			"",
			"",
			0,
		},
		{
			"sddd",
			"",
			0,
		},
		{
			"",
			"dffff",
			-1,
		},
		{
			"qwertyuiop",
			"qwert",
			0,
		},
		{
			"qwertyuiop",
			"wert",
			1,
		},
		{
			"qwertyuiop",
			"ertyuiop",
			2,
		},
		{
			"qwertyuiop",
			"rtyuiop",
			3,
		},
		{
			"aabbccddaabcbccdd",
			"bbccdd",
			2,
		},
	}

	for _, test := range tests {
		fmt.Println(strStr(test.haystack, test.needle) == test.index)
	}
}
