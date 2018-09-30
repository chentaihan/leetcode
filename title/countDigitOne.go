package title

import "fmt"

/*
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
示例:
输入: 13
输出: 6
解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。
*/

func countDigitOne(n int) int {
	count := 1
	for i := 10; i <= n; i++ {
		num := i
		for num > 0 {
			left := num % 10
			if left == 1 {
				count++
			}
			num /= 10
		}
	}
	return count
}

func TestCountDigitOne() {
	tests := []struct {
		sour int
		dest int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			5,
			1,
		},
		{
			10,
			2,
		},
		{
			11,
			4,
		},
		{
			13,
			6,
		},
	}
	for _, test := range tests {
		fmt.Println(countDigitOne(test.sour) == test.dest)
	}
}
