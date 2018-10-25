package mathTag

import "fmt"

/**
171. Excel表列序号
 */

func titleToNumber(s string) int {
	num := 0
	for _, c := range s {
		num = num*26 + int(c-'A'+1)
	}
	return num
}

func TesttitleToNumber() {
	tests := []struct {
		s   string
		num int
	}{
		{
			"A",
			1,
		},
		{
			"B",
			2,
		},
		{
			"Y",
			25,
		},
		{
			"Z",
			26,
		},
		{
			"AA",
			27,
		},
		{
			"AB",
			28,
		},
		{
			"ZY",
			701,
		},
	}
	for _, test := range tests {
		num := titleToNumber(test.s)
		fmt.Println(num == test.num)
	}
}
