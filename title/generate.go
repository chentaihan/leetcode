package title

import "fmt"

/**
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4, 6, 4,1]
[1,5,10,10,5,1]
[1,6,15,20,15,6,1]
]
https://leetcode-cn.com/problems/pascals-triangle/description/
*/

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	buffer := make([][]int, numRows)
	buffer[0] = []int{1}
	for rowIndex := 1; rowIndex < numRows; rowIndex++ {
		row := make([]int, rowIndex+1)
		row[0], row[rowIndex] = 1, 1
		for columnIndex := 1; columnIndex < rowIndex; columnIndex++ {
			row[columnIndex] = buffer[rowIndex-1][columnIndex-1] + buffer[rowIndex-1][columnIndex]
		}
		buffer[rowIndex] = row
	}
	return buffer
}

func generateEqual(a, b [][]int) bool {
	if a == nil || b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestGenerate() {
	tests := []struct {
		numRows int
		result  [][]int
	}{
		{
			0,
			nil,
		},
		{
			1,
			[][]int{{1}},
		},
		{
			2,
			[][]int{{1}, {1, 1}},
		},
		{
			3,
			[][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			4,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			6,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}},
		},
		{
			7,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1},{1,6,15,20,15,6,1}},
		},
	}
	for _, test := range tests {
		fmt.Println(generateEqual(generate(test.numRows), test.result))
	}
}
