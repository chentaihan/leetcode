package title

import "fmt"

func reverseString(s string) string {
	buffer := make([]rune, len(s))
	index := len(s) - 1
	for _, c := range s {
		buffer[index] = c
		index--
	}
	return string(buffer)
}

func TestReverseString() {
	tests := []struct {
		sour string
		dest string
	}{
		{
			"",
			"",
		},
		{
			"a",
			"a",
		},
		{
			"abcd",
			"dcba",
		},
		{
			"abcdf",
			"fdcba",
		},
	}
	for _, test := range tests {
		fmt.Println(reverseString(test.sour) == test.dest)
	}
}
