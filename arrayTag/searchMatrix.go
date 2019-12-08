package arrayTag

import "fmt"

/**
74. 搜索二维矩阵
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：就是一个变形的二分查找
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start, middle, end := 0, 0, row*col-1
	for start <= end {
		middle = start + (end-start)/2
		i := middle / col
		j := middle % col
		if i == row {
			return false
		}
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return false
}

func TestsearchMatrix() {
	tests := []struct {
		matrix [][]int
		target int
		result bool
	}{
		{
			[][]int{{1, 1}},
			2,
			false,
		},
		{
			[][]int{{1, 3}},
			3,
			true,
		},
		{
			[][]int{{1, 3}, {4, 6}},
			6,
			true,
		},
		{
			[][]int{{1, 3}, {4, 6}, {7, 9}},
			7,
			true,
		},
		{
			[][]int{{1, 3}, {4, 6}, {7, 9}},
			1,
			true,
		},
		{
			[][]int{{1, 3}, {4, 6}, {7, 9}},
			9,
			true,
		},
	}
	for _,test := range tests {
		result := searchMatrix(test.matrix, test.target)
		fmt.Println(result == test.result)
	}
}
