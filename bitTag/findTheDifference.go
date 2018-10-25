package bitTag

import "fmt"

/**
389. 找不同
分析：两个字符串字母相同，只是新增了一个字符，找出这个字符，异或运算
 */

func findTheDifference(s string, t string) byte {
	result := t[0]
	for i := 1; i < len(t); i++ {
		result ^= t[i]
	}
	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}
	return result
}

func TestfindTheDifference() {
	tests := []struct {
		s string
		t string
		r byte
	}{
		{
			"qwerty",
			"qwertyu",
			'u',
		},
		{
			"qwerty",
			"qweqrty",
			'q',
		},
		{
			"qwerty",
			"qwerrty",
			'r',
		},
		{
			"qweqrty",
			"qwerqtyq",
			'q',
		},{
			"qwteqrty",
			"qwetrqtyt",
			't',
		},
	}
	for _, test := range tests {
		fmt.Println(findTheDifference(test.t, test.s) == test.r)
	}
}
