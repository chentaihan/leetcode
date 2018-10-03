package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/*
73. 矩阵置零
 */

func setZeroes(matrix [][]int) {
	yLen := len(matrix)
	if yLen== 0 {
		return
	}
	xLen := len(matrix[0])
	buffer := make([]point, 0, len(matrix))
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if matrix[y][x] == 0 {
				buffer = append(buffer, point{x, y})
			}
		}
	}
	for _, p := range buffer {
		for x := 0; x < xLen; x++ {
			matrix[p.y][x] = 0
		}
		for y := 0; y < yLen; y++ {
			matrix[y][p.x] = 0
		}
	}
}

func TestsetZeroes() {
	tests := []struct{
		matrix [][]int
		result [][]int
	}{
		{
			[][]int{{1,1,1},{1,0,1},{1,1,1}},
			[][]int{{1,0,1},{0,0,0},{1,0,1}},
		},
		{
			[][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}},
			[][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}},
		},
	}
	for _,test := range tests{
		setZeroes(test.matrix)
		fmt.Println(common.IntEqualEx(test.matrix, test.result))
	}
}
