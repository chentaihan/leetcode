package stringTag

import "fmt"

/**
557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 */

func reverseWords(s string) string {
	buffer := make([]byte, len(s))
	start, index := 0, 0
	for i, c := range s {
		if c == ' ' {
			for j := i - 1; j >= start; j-- {
				buffer[index] = s[j]
				index++
			}
			start = i + 1
			buffer[index] = s[i]
			index++
		}
	}
	for i := len(s) - 1; i >= start; i-- {
		buffer[index] = s[i]
		index++
	}
	return string(buffer)
}

func TestreverseWords() {
	tests := []struct {
		s string
		r string
	}{
		{
			"a",
			"a",
		},
		{
			"ab",
			"ba",
		},
		{
			"abc",
			"cba",
		},
		{
			" abc",
			" cba",
		},
		{
			"abc ",
			"cba ",
		},
		{
			" abc ",
			" cba ",
		},
		{
			"abc def ghijk lm n",
			"cba fed kjihg ml n",
		},
		{
			"abc def ghijk lm n ",
			"cba fed kjihg ml n ",
		},
		{
			" abc def ghijk lm n",
			" cba fed kjihg ml n",
		},
		{
			" abc  def  ghijk  lm n ",
			" cba  fed  kjihg  ml n ",
		},
		{
			" abc  def zxcv ghijk  lm n ",
			" cba  fed vcxz kjihg  ml n ",
		},
	}
	for _, test := range tests {
		fmt.Println(reverseWords(test.s) == test.r)
	}
}
