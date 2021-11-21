package backtrackTag

import "fmt"

/**
980. 不同路径 III
在二维网格 grid 上，有 4 种类型的方格：

1 表示起始方格。且只有一个起始方格。
2 表示结束方格，且只有一个结束方格。
0 表示我们可以走过的空方格。
-1 表示我们无法跨越的障碍。
返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。

每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。

 

示例 1：

输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
输出：2
解释：我们有以下两条路径：
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
示例 2：

输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
输出：4
解释：我们有以下四条路径：
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
示例 3：

输入：[[0,1],[2,0]]
输出：0
解释：
没有一条路能完全穿过每一个空的方格一次。
请注意，起始和结束方格可以位于网格中的任意位置。

提示：

1 <= grid.length * grid[0].length <= 20

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func uniquePathsIII(grid [][]int) int {
	row, col, zeroCount := 0, 0, 1
	tag := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		tag[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				zeroCount++
				continue
			}
			if grid[i][j] == 1 {
				row = i
				col = j
			}
		}
	}
	result := 0
	_uniquePathsIII(grid, row, col, tag, &result, zeroCount)
	return result
}

func _uniquePathsIII(grid [][]int, row, col int, tag [][]bool, result *int, zeroCount int) {
	if grid[row][col] == 2 && 0 == zeroCount {
		*result++
		return
	}
	if tag[row][col] == true || grid[row][col] == -1 || grid[row][col] == 2 {
		return
	}

	tag[row][col] = true
	if row >= 1 { //上
		_uniquePathsIII(grid, row-1, col, tag, result, zeroCount-1)
	}
	if row < len(grid)-1 { //下
		_uniquePathsIII(grid, row+1, col, tag, result, zeroCount-1)
	}
	if col >= 1 { //左
		_uniquePathsIII(grid, row, col-1, tag, result, zeroCount-1)
	}
	if col < len(grid[0])-1 { //右
		_uniquePathsIII(grid, row, col+1, tag, result, zeroCount-1)
	}
	tag[row][col] = false
}

func TestUniquePathsIII() {
	tests := []struct {
		grid [][]int
		res  int
	}{
		{
			[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}},
			2,
		},
		{
			[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}},
			4,
		},
		{

			[][]int{{0, 1}, {2, 0}},
			0,
		},
	}
	for _, test := range tests {
		res := uniquePathsIII(test.grid)
		fmt.Println(res == test.res)
	}
}
