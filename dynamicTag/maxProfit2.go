package dynamicTag

import "fmt"

/**
122. 买卖股票的最佳时机 II
 */
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	sum := 0
	curVal := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			continue
		}
		sum += prices[i-1] - curVal
		curVal = prices[i]
	}

	sum += prices[len(prices)-1] - curVal
	return sum
}

func TestmaxProfit2() {
	tests := []struct {
		nums []int
		sum  int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		}, {
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			8,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
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
			10,
		},
		{
			[]int{-7, -2, -3, -1, -4, -5, -6},
			7,
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
			5,
		},
		{
			[]int{-2, 1, 0, 3},
			6,
		},
		{
			[]int{3, -2, 3, -2, 0, 2},
			9,
		},
		{
			[]int{3, -2, 3, -2, 0, 2, 0, 1},
			10,
		},
	}
	for _, test := range tests {
		fmt.Println(maxProfit2(test.nums) == test.sum)
	}
}
