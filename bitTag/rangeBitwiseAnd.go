package bitTag

import "fmt"

/**
201. 数字范围按位与
 */

func rangeBitwiseAnd(m int, n int) int {
	sum := m
	for m++; m <= n; m++ {
		sum &= m
	}
	return sum
}

func TestrangeBitwiseAnd() {
	tests := []struct {
		m int
		n int
		r int
	}{
		{
			4,
			7,
			4,
		},
	}
	for _, test := range tests {
		result := rangeBitwiseAnd(test.m, test.n)
		fmt.Println(result == test.r)
	}
}
