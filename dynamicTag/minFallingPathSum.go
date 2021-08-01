package dynamicTag

import "fmt"

/**
931. 下降路径最小和
给定一个方形整数数组 A，我们想要得到通过 A 的下降路径的最小和。

下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列。

示例：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：12
解释：
可能的下降路径有：
[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
和最小的下降路径是 [1,4,7]，所以答案是 12。

提示：

1 <= A.length == A[0].length <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-falling-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：当前格子的最小值 = min(上左值，上中值，上右值) + 当前值
最终最小值 = 结果数组最小值
*/

func minFallingPathSum(A [][]int) int {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}
	result := A[0]
	sumList := make([]int, len(A[0]))
	for row := 1; row < len(A); row++ {
		for col := 0; col < len(A[0]); col++ {
			min := result[col]
			if col > 0 && result[col-1] < min {
				min = result[col-1]
			}
			if col < len(A[0])-1 && result[col+1] < min {
				min = result[col+1]
			}
			sumList[col] = min + A[row][col]
		}
		result, sumList = sumList, result
	}
	min := result[0]
	for i := 1; i < len(result); i++ {
		if result[i] < min {
			min = result[i]
		}
	}
	return min
}

func minFallingPathSumEx(A [][]int) int {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}

	for row := 1; row < len(A); row++ {
		for col := 0; col < len(A[0]); col++ {
			min := A[row-1][col]
			if col > 0 && A[row-1][col-1] < min {
				min = A[row-1][col-1]
			}
			if col < len(A[0])-1 && A[row-1][col+1] < min {
				min = A[row-1][col+1]
			}
			A[row][col] += min
		}
	}
	sum := A[len(A)-1]
	min := sum[0]
	for i := 1; i < len(sum); i++ {
		if sum[i] < min {
			min = sum[i]
		}
	}
	return min
}

func TestminFallingPathSum() {
	tests := []struct {
		A   [][]int
		min int
	}{
		{
			[][]int{{10, -98, 44}, {-20, 65, 34}, {-100, -1, 74}},
			-218,
		},
		{
			[][]int{{1, 4, 9}, {7, 8, 2}, {3, 5, 2}},
			8,
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			12,
		},
		{
			[][]int{{1, 4, 9}, {7, 8, 2}, {3, 5, 6}},
			11,
		},
		{
			[][]int{{1, 4, 9}, {7, 2, 5}, {3, 3, 6}},
			6,
		},
		{
			[][]int{{3, 4, 1}, {7, 2, 5}, {3, 4, 6}},
			6,
		},
		{
			[][]int{{-19, 57}, {-40, -5}},
			-59,
		},
	}

	for _, test := range tests {
		result := minFallingPathSumEx(test.A)
		fmt.Println(result == test.min)
	}
}
