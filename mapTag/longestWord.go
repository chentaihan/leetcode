package mapTag

import "fmt"

/**
720. 词典中最长的单词
*/

func longestWord(words []string) string {
	m := make(map[string]struct{}, len(words))
	for _, word := range words {
		m[word] = struct{}{}
	}
	result := ""
	for _, word := range words {
		if len(word) >= len(result) {
			w := word
			for len(w) > 0 {
				if _, exist := m[w]; exist {
					if len(w) == 1 {
						if len(word) > len(result) || result > word {
							result = word
						}
						break
					}
					w = w[:len(w)-1]
				} else {
					break
				}
			}
		}

	}
	return result
}

func TestlongestWord() {
	tests := []struct {
		words  []string
		result string
	}{
		{
			[]string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"},
			"yodn",
		},
		{
			[]string{"w", "wo", "wor", "worl", "world"},
			"world",
		},
		{
			[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"},
			"apple",
		},
	}

	for _, test := range tests {
		result := longestWord(test.words)
		fmt.Println(result == test.result)
	}
}
