package mapTag

import (
	"fmt"
	"math"
)

/**
137. 只出现一次的数字 II
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
 */

func singleNumberEx(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	buffer := [64]int{}
	for _, num := range nums {
		for i := uint(0); i < 63; i++ {
			if 1<<i&num > 0 {
				buffer[i]++
			}
		}
		//考虑负数的情况
		if num < 0 {
			buffer[63]++
		}
	}

	result := 0
	for i := uint(0); i < 63; i++ {
		if buffer[i]%3 > 0 {
			result += 1 << i
		}
	}
	if buffer[63] % 3 > 0 {
		return result - math.MaxInt64 -1
	}
	return result
}

func TestsingleNumberEx() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 1, 1, 4},
			4,
		},
		{
			[]int{4},
			4,
		},
		{
			[]int{math.MaxInt64, math.MaxInt64, math.MaxInt64, 4, 1, 1, 1},
			4,
		},
		{
			[]int{math.MaxInt64 - 1, math.MaxInt64 - 1, math.MaxInt64 - 1, math.MaxInt64 - 3, 1, 1, 1},
			math.MaxInt64 - 3,
		},
		{
			[]int{1, 1, 1, 6, 6, 6, 4},
			4,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 4},
			4,
		},
		{
			[]int{3, 15, 3, 15, 3, 15, 4},
			4,
		},
		{
			[]int{23, 15, 23, 15, 23, 15, 4, 1, 1, 1},
			4,
		},
		{
			[]int{7, 7, 7, 2000, 1000, 2000, 40, 2000, 40, 40},
			1000,
		},
		{
			[]int{7, 7, 7, 2000, 2000, 1000, 2000, 40, 40, 1111, 1111, 40, 2345, 2345, 1111, 2345},
			1000,
		},
		{
			[]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2},
			-4,
		},
		{
			[]int{-2, -2, 1, 1, -3, 1, -3, -3, -5, -2},
			-5,
		},
		{
			[]int{-2, -2, 1, 1, -3, 1, -3, -3, -7, -2},
			-7,
		},
		{
			[]int{-2, -2, 1, 1, -3, 1, -3, -3, 1000, -2},
			1000,
		},
		{
			[]int{2, 2, 1, 1, -3, 1, -3, -3, 1000, 2},
			1000,
		},
	}
	for _, test := range tests {
		fmt.Println(singleNumberEx(test.nums) == test.result)
	}
}
