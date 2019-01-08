package arrayTag

/**
867. 转置矩阵
给定一个矩阵 A， 返回 A 的转置矩阵。

矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

示例 1：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
 */

func transpose(A [][]int) [][]int {
	rows := len(A)
	if rows == 0 {
		return [][]int{}
	}
	columns := len(A[0])
	buffer := make([][]int, columns)
	for i := 0; i < columns; i++ {
		buffer[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			buffer[j][i] = A[i][j]
		}
	}
	return buffer
}