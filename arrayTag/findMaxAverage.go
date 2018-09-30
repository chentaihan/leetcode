package arrayTag

import "fmt"

/**
643. 子数组最大平均数 I
 */

func findMaxAverage(nums []int, k int) int {
	sum, i := 0, 0
	for ; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for ; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum / k
}

func TestfindMaxAverage() {
	tests := []struct {
		nums   []int
		k      int
		result int
	}{
		{
			[]int{1, 2, 3, 4},
			4,
			10,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			30,
		},
		{
			[]int{11, 12, 13, 14, 5, 6, 7, 8, 9},
			4,
			50,
		},
		{
			[]int{4, 12, 13, 14, 5, 6, 7, 8, 9},
			4,
			44,
		},
		{
			[]int{4, 12, 13, 14, 5, 6, 7, 68, 9},
			4,
			90,
		},
	}
	for _, test := range tests {
		fmt.Println(findMaxAverage(test.nums, test.k) == test.result)
	}
}
