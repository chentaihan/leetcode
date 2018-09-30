package dynamicTag

import "fmt"

/**
121. 买卖股票的最佳时机
 */

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	nums := make([]int, len(prices))
	for i := 1; i < len(prices); i++{
		nums[i] = prices[i] - prices[i-1]
	}
	return maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		}else{
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func TestmaxProfit(){
	tests := []struct {
		nums []int
		sum  int
	}{
		{
			[]int{7,1,5,3,6,4},
			5,
		},{
			[]int{7,6,4,3,1},
			0,
		},
		{
			[]int{1,2,3,4,5,6,7,8,9},
			8,
		},
		{
			[]int{1,2,3,4,5,6,7,8,9,10},
			9,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 0},
			5,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 0},
			6,
		},
		{
			[]int{0, 0, 1, 2, 3, 4, 5, 6, 0, 0},
			6,
		},
		{
			[]int{0, -1, -2, -3, -4, -5, -6},
			0,
		},
		{
			[]int{-1, -2, -3, -4, -5, -6},
			0,
		},
		{
			[]int{-7, -1, -2, -3, -4, -5, -6},
			6,
		},
		{
			[]int{-7, -2, -3, -4, -5, -6, -1},
			6,
		},
		{
			[]int{-7, -2, -3, -1, -4, -5, -6},
			6,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{-1},
			0,
		},
		{
			[]int{0, 0, 0, 0, 0, 0},
			0,
		},
		{
			[]int{1, 0, -1},
			0,
		},
		{
			[]int{-1, 0, 1},
			2,
		},
		{
			[]int{2, -1, 0, 1},
			2,
		},
		{
			[]int{2, -1, 0, 2},
			3,
		},
		{
			[]int{-2, 1, 0, 2},
			4,
		},
		{
			[]int{-2, 1, 0, 3},
			5,
		},
		{
			[]int{3, -2, 3, -2, 0, 2},
			5,
		},
	}
	for _, test := range tests {
		fmt.Println(maxProfit(test.nums) == test.sum)
	}
}