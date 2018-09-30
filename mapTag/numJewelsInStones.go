package mapTag

import "fmt"

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}
	m := make(map[rune]bool, len(J))
	for _, c := range J {
		m[c] = true
	}

	count := 0
	for _, c := range S {
		if m[c] {
			count++
		}
	}
	return count
}

func TestNumJewelsInStones() {
	tests := []struct {
		J     string
		S     string
		count int
	}{
		{
			"",
			"",
			0,
		},
		{
			"aA",
			"aAAdsjkdsd",
			3,
		},
		{
			"aA",
			"AAdsjkdsd",
			2,
		},
		{
			"a",
			"AAdsjkdsd",
			0,
		},
		{
			"asdfghjkl",
			"asdfghjkl",
			9,
		},
		{
			"asdfghjkl",
			"asdfghjklasdfghjkl",
			18,
		},
		{
			"asdfghjkl",
			"asdfghjklasasdfghjklasdfghjkldfghjkl",
			36,
		},
	}
	for _, test := range tests {
		fmt.Println(numJewelsInStones(test.J, test.S) == test.count)
	}
}
