package dynamicTag

import "fmt"

/**
63. 不同路径 II
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n <= 0 {
		return 0
	}
	m := len(obstacleGrid[0])
	if m <= 0 {
		return 0
	}
	buffer := make([][]int, n)
	buffer[0] = make([]int, m)
	if obstacleGrid[0][0] == 0 {
		buffer[0][0] = 1
	}
	for i := 1; i < n; i++ {
		buffer[i] = make([]int, m)
		if obstacleGrid[i][0] == 1 {
			buffer[i][0] = 0
		} else {
			buffer[i][0] = buffer[i-1][0]
		}
	}

	for j := 1; j < m; j++ {
		if obstacleGrid[0][j] == 1 {
			buffer[0][j] = 0
		} else {
			buffer[0][j] = buffer[0][j-1]
		}
		for i := 1; i < n; i++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				buffer[i][j] = buffer[i-1][j] + buffer[i][j-1]
			}
		}
	}

	return buffer[n-1][m-1]
}

func TestuniquePathsWithObstacles() {
	tests := []struct {
		obstacleGrid [][]int
		result       int
	}{
		{
			[][]int{{0, 0}},
			1,
		},
		{
			[][]int{{1, 0}},
			0,
		},
	}

	for _, test := range tests {
		fmt.Println(uniquePathsWithObstacles(test.obstacleGrid) == test.result)
	}
}
