package arrayTag

import "fmt"

/**
1351. 统计有序矩阵中的负数
给你一个 m * n 的矩阵 grid，矩阵中的元素无论是按行还是按列，都以非递增顺序排列。 请你统计并返回 grid 中 负数 的数目。

示例 1：

输入：grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
输出：8
解释：矩阵中共有 8 个负数。
示例 2：

输入：grid = [[3,2],[1,0]]
输出：0


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 100
-100 <= grid[i][j] <= 100


进阶：你可以设计一个时间复杂度为 O(n + m) 的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countNegatives(grid [][]int) int {
	start, total := 0, 0
	for i := len(grid) - 1; i >= 0; i-- {
		start = countNegativesIndex(grid[i], start)
		total += len(grid[i]) - start
	}
	return total
}

func countNegativesIndex(array []int, start int) int {
	end := len(array) - 1
	for start <= end {
		middle := start + (end-start)/2
		if array[middle] < 0 {
			if middle > 0 && array[middle-1] >= 0 {
				return middle
			}
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return start
}

func TestcountNegatives() {
	tests := []struct {
		grid  [][]int
		count int
	}{
		{
			[][]int{{4, 3, 2}, {3, 2, -1}, {1, 1, -2}, {-1, -1, -3}},
			5,
		},
		{
			[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			8,
		},

		{
			[][]int{{4, 3, -1}, {3, 2, -2}},
			2,
		},
		{
			[][]int{{4, 3, 2}, {3, 2, 0}},
			0,
		},
		{
			[][]int{{4, 3, 2}, {3, 2, -1}},
			1,
		},
	}
	for _, test := range tests {
		count := countNegatives(test.grid)
		fmt.Println(count == test.count)
	}
}
