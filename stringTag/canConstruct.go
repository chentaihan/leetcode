package stringTag

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make(map[rune]int, len(magazine))
	for _, c := range magazine {
		m[c]++
	}
	for _, c := range ransomNote {
		count := m[c]
		if count <= 0 {
			return false
		}
		m[c]--
	}
	return true
}

func TestCanConstruct() {
	tests := []struct{
		ransomNote string
		magazine string
		result bool
	}{
		{
			"",
			"",
			true,
		},
		{
			"",
			"dfgh",
			true,
		},
		{
			"sdf",
			"",
			false,
		},
		{
			"aaaaaaaaa",
			"zxcvbnm",
			false,
		},
		{
			"aaaaaaaaa",
			"azxcvbnm",
			false,
		},
		{
			"aaaaaaaaa",
			"bbbbbbbbbbbb",
			false,
		},
		{
			"zxcvbnm",
			"mnbvcxz",
			true,
		},
		{
			"zxcvbnmzxcvbnmzxcvbnmzxcvbnmzxcvbnmzxcvbnmzxcvbnmzxcvbnm",
			"mnbvcxz",
			false,
		},
		{
			"zzxcvbnm",
			"mnbvcxz",
			false,
		},
		{
			"zxcvbnm",
			"mnbvcxzdfdccv",
			true,
		},
		{
			"zxcvbnmzxcvbnm",
			"mnbvczxcvbnmxz",
			true,
		},
	}
	for _,test := range tests{
		fmt.Println(canConstruct(test.ransomNote, test.magazine) == test.result)
	}
}
