package dynamicTag

import "fmt"

/**
70. 爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
 */

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, next := 1, 2
	ret := 0
	for i := 3; i <= n; i++ {
		ret = prev + next
		prev = next
		next = ret
	}
	return ret
}

func TestclimbStairs() {
	tests := []struct {
		n   int
		ret int
	}{
		{
			1, 1,
		},
		{
			2, 2,
		},
		{
			3, 3,
		},
		{
			4, 5,
		},
		{
			5, 8,
		},
		{
			6, 13,
		},
		{
			7, 21,
		},
	}

	for _, test := range tests {
		fmt.Println(climbStairs(test.n) == test.ret)
	}
}
