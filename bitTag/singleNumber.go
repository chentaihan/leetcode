package bitTag

import "fmt"

/**
136. 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
 */

func singleNumber1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := len(nums)-1
	result := nums[i]
	i--
	for i >= 8 {
		result ^= nums[i];i--
		result ^= nums[i];i--
		result ^= nums[i];i--
		result ^= nums[i];i--
		result ^= nums[i];i--
		result ^= nums[i];i--
		result ^= nums[i];i--
		result ^= nums[i];i--
	}
	switch i {
	case 7:
		result ^= nums[i]; i--; fallthrough
	case 6:
		result ^= nums[i]; i--; fallthrough
	case 5:
		result ^= nums[i]; i--; fallthrough
	case 4:
		result ^= nums[i]; i--; fallthrough
	case 3:
		result ^= nums[i]; i--; fallthrough
	case 2:
		result ^= nums[i]; i--; fallthrough
	case 1:
		result ^= nums[i]; i--; fallthrough
	case 0:
		result ^= nums[i]
	}
	return result
}

func TestsingleNumber() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 1, 2, 2, 4},
			4,
		},
		{
			[]int{7, 7, 2000, 1000, 2000, 40, 40},
			1000,
		},
		{
			[]int{7, 7, 2000, 1000, 2000, 40, 40, 1111, 2345, 1111, 2345},
			1000,
		},
	}
	for _, test := range tests {
		fmt.Println(singleNumber(test.nums) == test.result)
	}
}
