package mapTag

import "fmt"

/**
463. 岛屿的周长
给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

示例 :

[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

答案: 16
https://leetcode-cn.com/problems/island-perimeter/description/
 */

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) ==0{
		return 0
	}
	yLen := len(grid)
	xLen := len(grid[0])
	perimeter := 0
	//计算第一行
	if grid[0][0] == 1 {
		perimeter = 4
	}
	for i := 1; i < xLen; i++ {
		if grid[0][i] == 1 {
			if grid[0][i-1] == 1 {
				perimeter += 2
			}else{
				perimeter += 4
			}
		}
	}
	//其他行
	for i := 1; i < yLen;i++{
		if grid[i][0] == 1 {
			if grid[i-1][0] == 1 {
				perimeter += 2
			}else{
				perimeter += 4
			}
		}

		for j := 1; j < xLen; j++ {
			if grid[i][j] == 0 {
				continue
			}
			count := 4
			if grid[i-1][j] == 1 {
				count -= 2
			}
			if grid[i][j-1] == 1 {
				count -= 2
			}
			perimeter += count
		}
	}
	return perimeter
}

func TestislandPerimeter(){
	tests := []struct{
		grid [][]int
		perimeter int
	}{
		{
			[][]int{{1,0}},
			4,
		},
		{
			[][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}},
			16,
		},
	}
	for _,test := range tests{
		fmt.Println(islandPerimeter(test.grid) == test.perimeter)
	}
}