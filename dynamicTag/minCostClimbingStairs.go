package dynamicTag

import "fmt"

/**
746. 使用最小花费爬楼梯
数组的每个索引做为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:

输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
 示例 2:

输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
 */

func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 2 {
		return 0
	}
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func TestminCostClimbingStairs() {
	tests := []struct {
		cost []int
		ret  int
	}{
		{
			[]int{10, 15, 20},
			15,
		},
		{
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 1},
			4,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 1, 1},
			4,
		},
		{
			[]int{1, 2, 1, 2, 1, 2, 1, 1, 2},
			5,
		},
	}

	for _, test := range tests {
		fmt.Println(minCostClimbingStairs(test.cost) == test.ret)
	}
}
