package dynamicTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

//剑指 Offer II 107. 矩阵中的距离
//https://leetcode-cn.com/problems/2bCMpM/
func updateMatrix(mat [][]int) [][]int {
	row, col := len(mat), len(mat[0])
	queue := make([][2]int, 0, row*col)
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				result[i][j] = -1
			}
		}
	}
	flag := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		item := queue[0]
		value := result[item[0]][item[1]]
		for i := 0; i < len(flag); i++ {
			x, y := item[0]+flag[i][0], item[1]+flag[i][1]
			if x >= 0 && x < row && y >= 0 && y < col {
				if result[x][y] == -1 && mat[x][y] != 0 {
					result[x][y] = value + 1
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		queue = queue[1:]
	}
	return result
}

func TestupdateMatrix() {
	tests := []struct {
		mat    [][]int
		result [][]int
	}{
		{
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
		{
			[][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			[][]int{
				{0, 1, 2},
				{1, 2, 3},
				{2, 3, 4},
			},
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{2, 1, 2},
				{1, 0, 1},
				{2, 1, 2},
			},
		},
		{
			[][]int{
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 1, 1},
			},
			[][]int{
				{2, 1, 2, 3},
				{1, 0, 1, 2},
				{2, 1, 2, 3},
			},
		},
		{
			[][]int{
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			[][]int{
				{2, 1, 2, 3},
				{1, 0, 1, 2},
				{2, 1, 2, 3},
				{3, 2, 3, 4},
			},
		},
	}
	for index, test := range tests {
		result := updateMatrix(test.mat)
		isOk := true
		for i := 0; i < len(result); i++ {
			if !common.IntEqual(result[i], test.result[i]) {
				isOk = false
				break
			}
		}
		fmt.Println(index, isOk)
	}
}
