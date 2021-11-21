package backtrackTag

/**
剑指 Offer 12. 矩阵中的路径
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
 

提示：

1 <= board.length <= 200
1 <= board[i].length <= 200
board 和 word 仅由大小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func exist(board [][]byte, word string) bool {
	tag := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		tag[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if _exist(board, tag, i, j, 0, word) {
					return true
				}
			}
		}
	}
	return false
}

func _exist(board [][]byte, tag [][]bool, row, col int, index int, word string) bool {
	if row <= -1 || col <= -1 || row >= len(board) || col >= len(board[0]) || index >= len(word) {
		return false
	}
	if tag[row][col] || board[row][col] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	tag[row][col] = true
	ok := _exist(board, tag, row-1, col, index+1, word) || _exist(board, tag, row+1, col, index+1, word) || _exist(board, tag, row, col-1, index+1, word) || _exist(board, tag, row, col+1, index+1, word)
	tag[row][col] = false
	return ok
}
