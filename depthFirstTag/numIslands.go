package depthFirstTag

/**
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numIslands(grid [][]byte) int {
	dp := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]bool, len(grid[0]))
	}
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if !dp[row][col] && grid[row][col] == '1' {
				count++
				_numIslands(grid, dp, row, col)
			}
		}
	}
	return count
}

func _numIslands(grid [][]byte, dp [][]bool, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' || dp[row][col] {
		return
	}
	dp[row][col] = true
	_numIslands(grid, dp, row-1, col) //上
	_numIslands(grid, dp, row+1, col) //下
	_numIslands(grid, dp, row, col-1) //左
	_numIslands(grid, dp, row, col+1) //右
}
