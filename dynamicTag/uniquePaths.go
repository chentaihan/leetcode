package dynamicTag

import "fmt"

/**
62. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？

例如，上图是一个7 x 3 的网格。有多少可能的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28
*/

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	buffer := make([][]int, n)
	for i := 0; i < n; i++ {
		buffer[i] = make([]int, m)
		buffer[i][0] = 1
	}

	for j := 1; j < m; j++ {
		buffer[0][j] = 1
		for i := 1; i < n; i++ {
			buffer[i][j] = buffer[i-1][j] + buffer[i][j-1]
		}
	}

	return buffer[n-1][m-1]
}

func uniquePathsEx(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	buf := make([][]int, m)
	for i := 0; i < m; i++ {
		buf[i] = make([]int, n)
		buf[i][0] = 1
	}

	for col :=1; col < n; col++ {
		buf[0][col] = 1
		for row := 1; row < m; row++ {
			buf[row][col] = buf[row-1][col] + buf[row][col-1]
		}
	}
	return buf[m-1][n-1]
}

func TestuniquePaths() {
	tests := []struct {
		m   int
		n   int
		ret int
	}{
		{
			3,
			2,
			3,
		},
		{
			2,
			2,
			2,
		},
		{
			1,
			2,
			1,
		},
		{
			3,
			3,
			6,
		},{
			4,
			4,
			20,
		},
		{
			7,
			3,
			28,
		},
	}

	for _, test := range tests {
		fmt.Println(uniquePathsEx(test.m, test.n) == test.ret)
	}
}
