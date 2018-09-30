package dynamicTag

import "fmt"

/**
最长公共连续字符串字符串
 */
func maxSubString(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	maxLen := 0
	buffer := make([][]int, len(str2))
	for i := 0; i < len(str2); i++ {
		buffer[i] = make([]int, len(str1))
		for j := 0; j < len(str1); j++ {
			if str1[j] == str2[i] {
				if i > 0 && j > 0 {
					buffer[i][j] = buffer[i-1][j-1] + 1
				} else {
					buffer[i][j] = 1
				}
				if buffer[i][j] > maxLen {
					maxLen = buffer[i][j]
				}
			}
		}
	}
	return maxLen
}

/**
最长公共连续字符串字符串（优化版）
 */
func maxSubString1(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	maxLen := 0
	prev := make([]int, len(str1))
	cur := make([]int, len(str1))
	for i := 0; i < len(str2); i++ {
		prev, cur = cur, prev
		for j := 0; j < len(str1); j++ {
			if str1[j] == str2[i] {
				if j > 0 {
					cur[j] = prev[j-1] + 1
				} else {
					cur[j] = 1
				}
				if cur[j] > maxLen {
					maxLen = cur[j]
				}
			}else{
				cur[j] = 0
			}
		}
	}
	return maxLen
}

func TestmaxSubString() {
	tests := []struct {
		str1   string
		str2   string
		result int
	}{
		{
			"1234567890",
			"5678",
			4,
		},
		{
			"1234567890",
			"dsd5678vf",
			4,
		},
		{
			"1234567890",
			"dsd5678v89f",
			4,
		},
		{
			"1234567890a",
			"dsd5678v7890af",
			5,
		},
		{
			"1234567890vf",
			"dsd5678vf",
			4,
		},
		{
			"1234567890",
			"11",
			1,
		},
		{
			"1234567890",
			"111111111111",
			1,
		},
		{
			"1234567890",
			"dsdsdfgvf",
			0,
		},
		{
			"1234567890",
			"1234567890",
			10,
		},
		{
			"1234567890x",
			"1234567890",
			10,
		},
		{
			"x1234567890",
			"1234567890",
			10,
		},
		{
			"1234567890",
			"x1234567890",
			10,
		},
		{
			"1234567890",
			"1234567890x",
			10,
		},
		{
			"a1234567890b",
			"asdc1234567890dddd",
			10,
		},
	}
	for _, test := range tests {
		fmt.Println(maxSubString1(test.str1, test.str2) == test.result)
		//fmt.Println(maxSubString(test.str1, test.str2) == test.result)
	}
}
