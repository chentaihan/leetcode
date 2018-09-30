package title

import "fmt"

/**
给定一个正整数 n，你可以做如下操作：

1. 如果 n 是偶数，则用 n / 2替换 n。
2. 如果 n 是奇数，则可以用 n + 1或n - 1替换 n。
n 变为 1 所需的最小替换次数是多少？

示例 1:

输入:
8

输出:
3

解释:
8 -> 4 -> 2 -> 1
示例 2:

输入:
7

输出:
4

解释:
7 -> 8 -> 4 -> 2 -> 1
或
7 -> 6 -> 3 -> 2 -> 1
https://leetcode-cn.com/problems/integer-replacement/description/

分析：看奇数前后两个偶数二进制中1的个数，1多的步数 >= 1少的步数
 */

func binary1Conut(x int) int {
	if x <= 0 {
		return 0
	}
	count := 0
	for x > 0 {
		count++
		x = x & (x - 1)
	}
	return count
}

func integerReplacement(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			count1 := binary1Conut(n + 1)
			count2 := binary1Conut(n - 1)
			if count1 >= count2 {
				n--
			} else {
				n++
			}
		}
		count++
	}
	return count
}

func TestIntegerReplacement() {
	tests := []struct {
		n     int
		count int
	}{
		{
			0,
			0,
		},
		{
			1,
			0,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			2,
		},
		{
			5,
			3,
		},
		{
			6,
			3,
		},
		{
			7,
			4,
		},
		{
			8,
			3,
		},
		{
			9,
			4,
		},
		{
			10,
			4,
		},
		{
			11,
			5,
		},
		{
			12,
			4,
		},
		{
			13,
			5,
		},
		{
			14,
			5,
		},
		{
			15,
			5,
		},
		{
			16,
			4,
		},
		{
			17,
			5,
		},
		{
			18,
			5,
		},
		{
			19,
			6,
		},
		{
			20,
			5,
		},
		{
			21,
			6,
		},
		{
			22,
			6,
		},
		{
			23,
			6,
		},
		{
			24,
			5,
		},
		{
			25,
			6,
		},
		{
			26,
			6,
		},
		{
			27,
			7,
		},
		{
			28,
			6,
		},
		{
			29,
			7,
		},
		{
			30,
			6,
		},
		{
			31,
			6,
		},
		{
			32,
			5,
		},
		{
			33,
			6,
		},
		{
			34,
			6,
		},
	}

	for _, test := range tests {
		fmt.Println(integerReplacement(test.n) == test.count)
	}
}
