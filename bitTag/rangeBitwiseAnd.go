package bitTag

import "fmt"

/**
201. 数字范围按位与
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

示例 1: 

输入: [5,7]
输出: 4
示例 2:

输入: [0,1]
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bitwise-and-of-numbers-range
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rangeBitwiseAnd(m int, n int) int {
	if m <= 0 || n / 2 >= m {
		return 0
	}
	first, second := 0, 1
	for first < 2147483648 {
		if first <= m && n < second {
			return first + rangeBitwiseAnd(m-first, n-first)
		}
		first = second
		second <<= 1
	}
	return 0
}

func TestrangeBitwiseAnd() {
	tests := []struct {
		m int
		n int
		r int
	}{
		{
			0,
			5,
			0,
		},
		{
			5,
			7,
			4,
		},
		{
			4,
			7,
			4,
		},
		{
			12,
			15,
			12,
		},
		{
			8,
			16,
			0,
		},
	}
	for _, test := range tests {
		result := rangeBitwiseAnd(test.m, test.n)
		fmt.Println(result == test.r)
	}
}
