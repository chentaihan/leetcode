package arrayTag

import "fmt"

/**
1281. 整数的各位积和之差
给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。



示例 1：

输入：n = 234
输出：15
解释：
各位数之积 = 2 * 3 * 4 = 24
各位数之和 = 2 + 3 + 4 = 9
结果 = 24 - 9 = 15
示例 2：

输入：n = 4421
输出：21
解释：
各位数之积 = 4 * 4 * 2 * 1 = 32
各位数之和 = 4 + 4 + 2 + 1 = 11
结果 = 32 - 11 = 21


提示：

1 <= n <= 10^5


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func subtractProductAndSum(n int) int {
	if n < 10 {
		return 0
	}
	total, sum := 1, 0
	for n > 0 {
		val := n % 10
		n /= 10
		total *= val
		sum += val
	}
	return total - sum
}

func TestSubtractProductAndSum() {
	tests := []struct {
		n   int
		res int
	}{
		{
			234, 15,
		},
		{
			4421, 21,
		},
		{
			11, -1,
		},
		{
			9, 0,
		},
		{
			10, -1,
		},
	}
	for _, test := range tests {
		res := subtractProductAndSum(test.n)
		if res == test.res {
			fmt.Println(true)
		} else {
			fmt.Println(res, test.res)
		}
	}
}
