package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
977. 有序数组的平方
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。



示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：变形的合并排序
*/

func sortedSquares(A []int) []int {
	index := 0
	for ; index < len(A); index++ {
		if A[index] >= 0 {
			break
		}
	}
	arr1 := A[:index]
	arr2 := A[index:]
	start, end := 0, len(arr1)-1
	for start <= end {
		arr1[start], arr1[end] = 0-arr1[end], 0-arr1[start]
		start++
		end--
	}
	index1, index2 := 0, 0
	index = 0
	result := make([]int, len(A))
	for index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1] < arr2[index2] {
			result[index] = arr1[index1] * arr1[index1]
			index1++
		} else {
			result[index] = arr2[index2] * arr2[index2]
			index2++
		}
		index++
	}
	for index1 < len(arr1) {
		result[index] = arr1[index1] * arr1[index1]
		index1++
		index++
	}
	for index2 < len(arr2) {
		result[index] = arr2[index2] * arr2[index2]
		index2++
		index++
	}
	return result
}

func TestsortedSquares() {
	tests := []struct {
		A      []int
		result []int
	}{
		{
			[]int{-2, 0},
			[]int{0, 4},
		},
	}

	for _, test := range tests {
		result := sortedSquares(test.A)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
