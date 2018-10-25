package mathTag

import "fmt"

/**
168. Excel表列名称
 */

func convertToTitle(n int) string {
	result := ""

	return result
}

func TestconvertToTitle() {
	tests := []struct {
		n int
		r string
	}{
		{
			1,
			"A",
		},
		{
			2,
			"B",
		},
		{
			25,
			"Y",
		},
		{
			26,
			"Z",
		},
		{
			27,
			"AA",
		},
		{
			28,
			"AB",
		},
		{
			52,
			"AZ",
		},
		{
			701,
			"ZY",
		},
		{
			703,
			"AAA",
		},
	}
	for _, test := range tests {
		r := convertToTitle(test.n)
		fmt.Println(r == test.r)
	}
}
