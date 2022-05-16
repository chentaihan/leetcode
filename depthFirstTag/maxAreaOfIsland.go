package depthFirstTag

/**
695. 岛屿的最大面积
给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。



示例 1：


输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
示例 2：

输入：grid = [[0,0,0,0,0,0,0,0]]
输出：0


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxAreaOfIsland(grid [][]int) int {
	dp := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]bool, len(grid[0]))
	}
	max := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if !dp[row][col] && grid[row][col] == 1 {
				count := 0
				_maxAreaOfIsland(grid, dp, row, col, &count)
				if count > max {
					max = count
				}
			}
		}
	}
	return max
}

func _maxAreaOfIsland(grid [][]int, dp [][]bool, row, col int, max *int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == 0 || dp[row][col] {
		return
	}
	dp[row][col] = true
	*max++
	_maxAreaOfIsland(grid, dp, row-1, col, max) //上
	_maxAreaOfIsland(grid, dp, row+1, col, max) //下
	_maxAreaOfIsland(grid, dp, row, col-1, max) //左
	_maxAreaOfIsland(grid, dp, row, col+1, max) //右
}
