package bitTag

import (
	"fmt"
)

/**
326. 3的幂
分析：一直除以3，最后必然等于1
 */

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

func TestisPowerOfThree() {
	tests := []struct {
		num    int
		result bool
	}{
		{
			0,
			false,
		},
		{
			1,
			true,
		},
		{
			3,
			true,
		},
		{
			6,
			false,
		},
		{
			9,
			true,
		},
		{
			18,
			false,
		},
		{
			27,
			true,
		},
		{
			81,
			true,
		},
		{
			243,
			true,
		},
		{
			324,
			false,
		},
		{
			246,
			false,
		},
	}
	for _, test := range tests {
		result := isPowerOfThree(test.num)
		fmt.Println(result == test.result)
	}
}
