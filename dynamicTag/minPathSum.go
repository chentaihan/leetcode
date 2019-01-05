package dynamicTag

import (
	"fmt"
)

/**
64. 最小路径和
https://leetcode-cn.com/problems/minimum-path-sum/description/
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
解析：先计算第一行和第一列到每个单元格的最小和，然后通过这些计算剩余每个单元最小值总和
当前单元格总和最小值=当前单元格值 + min(当前单元格上,当前单元格左)
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func TestminPathSum() {
	tests := []struct {
		grid [][]int
		ret  int
	}{
		{
			[][]int{
			},
			0,
		},
		{
			[][]int{
				{1},
			},
			1,
		},
		{
			[][]int{
				{1, 3},
				{5, 3},
			},
			7,
		},
		{
			[][]int{
				{1, 3, 1},
				{5, 3, 2},
				{1, 3, 1},
			},
			8,
		},
		{
			[][]int{
				{1, 3, 1},
				{1, 3, 2},
				{1, 3, 1},
			},
			7,
		},
		{
			[][]int{
				{1, 3, 1},
				{1, 0, 2},
				{1, 0, 1},
			},
			3,
		},
		{
			[][]int{
				{1, 3, 1, 2},
				{1, 0, 2, 4},
				{1, 0, 2, 2},
				{1, 0, 1, 1},
			},
			4,
		},
		{
			[][]int{
				{1, 3, 1, 2},
				{1, 0, 2, 4},
				{1, 0, 2, 1},
				{1, 0, 2, 1},
			},
			5,
		},
		{
			[][]int{
				{1, 3, 1, 2},
				{1, 0, 2, 4},
				{1, 2, 1, 1},
				{1, 0, 0, 1},
			},
			5,
		},
	}

	for _, test := range tests {
		fmt.Println(minPathSum(test.grid) == test.ret)
	}
}
