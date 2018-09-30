package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
88. 合并两个有序数组
 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	curIndex, index1, index2 := m+n-1, m-1, n-1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] >= nums2[index2] {
			nums1[curIndex] = nums1[index1]
			index1--
		} else {
			nums1[curIndex] = nums2[index2]
			index2--
		}
		curIndex--
	}
	for index1 >= 0 {
		nums1[curIndex] = nums1[index1]
		curIndex--
		index1--
	}
	for index2 >= 0 {
		nums1[curIndex] = nums2[index2]
		curIndex--
		index2--
	}
}

func TestMerge() {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{
			[]int{1, 0},
			1,
			[]int{1},
			1,
			[]int{1, 1},
		},
		{
			[]int{2, 0},
			1,
			[]int{1},
			1,
			[]int{1, 2},
		},
		{
			[]int{1, 0},
			1,
			[]int{2},
			1,
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3, 4, 0, 0, 0, 0},
			4,
			[]int{1, 2, 3, 4},
			4,
			[]int{1, 1, 2, 2, 3, 3, 4, 4},
		},
		{
			[]int{1, 2, 0, 0, 0, 0},
			2,
			[]int{1, 2, 3, 4},
			4,
			[]int{1, 1, 2, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 0, 0},
			4,
			[]int{1, 2},
			2,
			[]int{1, 1, 2, 2, 3, 4},
		},
	}
	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		fmt.Println(common.IntEqual(test.nums1, test.result))
	}
}
