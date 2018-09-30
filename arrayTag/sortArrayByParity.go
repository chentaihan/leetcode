package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
905. 按奇偶校验排序数组
 */

func sortArrayByParity(A []int) []int {
	start, end := 0, len(A)-1
	for start < end {
		for start < end && A[start]%2 == 0 {
			start++
		}
		for end >= 0 && A[end]%2 == 1 {
			end--
		}
		if start < end {
			A[start], A[end] = A[end], A[start]
		}
	}
	return A
}

func TestsortArrayByParity() {
	tests := []struct {
		A []int
		B []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{4, 2, 3, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 2, 3, 1, 5},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{4, 2, 6, 3, 1, 5},
		},
		{
			[]int{1, 2, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 2, 6, 3, 1, 5, 7},
		},
		{
			[]int{1, 2, 2, 3, 4, 8, 5, 6, 7},
			[]int{4, 2, 2, 6, 8, 3, 1, 5, 7},
		},
		{
			[]int{1, 1, 1, 1, 2, 2, 2, 2},
			[]int{2, 2, 2, 2, 1, 1, 1, 1},
		},
		{
			[]int{2, 2, 2, 2, 1, 1, 1, 1},
			[]int{2, 2, 2, 2, 1, 1, 1, 1},
		},
		{
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 1, 1, 1, 1},
		},
		{
			[]int{2, 2, 2, 2, 2},
			[]int{2, 2, 2, 2, 2},
		},
		{
			[]int{2, 2, 2, 2, 2, 2},
			[]int{2, 2, 2, 2, 2, 2},
		},
	}
	for _, test := range tests {
		B := sortArrayByParity(test.A)
		fmt.Println(common.IntEqualSort(test.B, B))
	}
}
