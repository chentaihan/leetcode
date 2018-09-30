package title

import "fmt"

func toLowerCase(str string) string {
	buffer := make([]rune, len(str))
	for index, c := range str {
		if c >= 'A' && c <= 'Z' {
			buffer[index] = c + 32
		} else {
			buffer[index] = c
		}
	}
	return string(buffer)
}

func TestToLowerCase() {
	tests := []struct {
		sour string
		dest string
	}{
		{
			"",
			"",
		},
		{
			"Hello",
			"hello",
		},
	}

	for _, test := range tests {
		fmt.Println(toLowerCase(test.sour) == test.dest)
	}
}
