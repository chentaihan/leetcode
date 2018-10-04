package bitTag

import "fmt"

/**
461. 汉明距离
 */

func hammingDistance(x int, y int) int {
	count := 0
	for x > 0 && y > 0 {
		if (x&1)^(y&1) == 1 {
			count++
		}
		x >>= 1
		y >>= 1
	}
	for ; x > 0; x = x & (x - 1) {
		count++
	}
	for ; y > 0; y = y & (y - 1) {
		count++
	}
	return count
}

func TesthammingDistance() {
	tests := []struct {
		x int
		y int
		r int
	}{
		{
			1,
			4,
			2,
		},
		{
			7,
			8,
			4,
		},
		{
			7,
			6,
			1,
		},
		{
			7,
			5,
			1,
		},
		{
			7,
			4,
			2,
		},
		{
			4,
			3,
			3,
		},
		{
			15,
			16,
			5,
		},
		{
			31,
			32,
			6,
		},
		{
			31,
			15,
			1,
		},
		{
			32,
			15,
			5,
		},
		{
			64,
			15,
			5,
		},
	}
	for _, test := range tests {
		r := hammingDistance(test.x, test.y)
		fmt.Println(r == test.r)
	}
}
