package mapTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

func findWords(words []string) []string {
	m1 := make(map[byte]bool, 10)
	for _, c := range "qwertyuiopQWERTYUIOP" {
		m1[byte(c)] = true
	}
	m2 := make(map[byte]bool, 9)
	for _, c := range "asdfghjklASDFGHJKL" {
		m2[byte(c)] = true
	}
	m3 := make(map[byte]bool, 7)
	for _, c := range "zxcvbnmZXCVBNM" {
		m3[byte(c)] = true
	}
	ret := []string{}
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if m1[word[0]] && checkWord(m1, word) {
			ret = append(ret, word)
			continue
		}
		if m2[word[0]] && checkWord(m2, word) {
			ret = append(ret, word)
			continue
		}
		if m3[word[0]] && checkWord(m3, word) {
			ret = append(ret, word)
		}
	}
	return ret
}

func checkWord(m map[byte]bool, word string) bool {
	for _, c := range word {
		if !m[byte(c)] {
			return false
		}
	}
	return true
}

func TestfindWords() {
	tests := []struct {
		words  []string
		result []string
	}{
		{
			[]string{"qwer", "asdf", "zxcvbcxz"},
			[]string{"qwer", "asdf", "zxcvbcxz"},
		},
		{
			[]string{"qwefr", "asdf", "zxcvbdcxz"},
			[]string{"asdf"},
		},
		{
			[]string{"qwer", "aqsdf", "zxcfvbcxz"},
			[]string{"qwer"},
		},
		{
			[]string{"qwser", "asduf", "zxcvbcxz"},
			[]string{"zxcvbcxz"},
		},
		{
			[]string{"Hello","Alaska","Dad","Peace"},
			[]string{"Alaska","Dad"},
		},

	}
	for _, test := range tests {
		fmt.Println( common.StringEqualEx(findWords(test.words), test.result))
	}
}
