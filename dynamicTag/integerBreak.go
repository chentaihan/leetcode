package dynamicTag

import "fmt"

/**
343. 整数拆分
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	max := 0
	for i := 2; ; i++ {
		val := n / i
		left := n - val*i
		sum := val
		for j := 1; j < i-left; j++ {
			sum *= val
		}
		val++
		for j := 0; j < left; j++ {
			sum *= val
		}
		if sum > max {
			max = sum
		}
		if val < 2 {
			break
		}
	}
	return max
}

func TestintegerBreak() {
	tests := []struct {
		n      int
		result int
	}{
		{
			10,
			36,
		},
		{
			2,
			1,
		},
		{
			4,
			4,
		},
		{
			3,
			2,
		},
		{
			5,
			6,
		},
		{
			6,
			9,
		},
		{
			7,
			12,
		},
		{
			8,
			18,
		},
		{
			9,
			27,
		},
		{
			11,
			54,
		},
		{
			16,
			324,
		},
	}

	for index,test := range tests {
		result := integerBreak(test.n)
		fmt.Println(index, result == test.result)
	}
}
