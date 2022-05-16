package arrayTag

import "fmt"

/**
1323. 6 和 9 组成的最大数字
给你一个仅由数字 6 和 9 组成的正整数 num。

你最多只能翻转一位数字，将 6 变成 9，或者把 9 变成 6 。

请返回你可以得到的最大数字。



示例 1：

输入：num = 9669
输出：9969
解释：
改变第一位数字可以得到 6669 。
改变第二位数字可以得到 9969 。
改变第三位数字可以得到 9699 。
改变第四位数字可以得到 9666 。
其中最大的数字是 9969 。
示例 2：

输入：num = 9996
输出：9999
解释：将最后一位从 6 变到 9，其结果 9999 是最大的数。
示例 3：

输入：num = 9999
输出：9999
解释：无需改变就已经是最大的数字了。


提示：

1 <= num <= 10^4
num 每一位上的数字都是 6 或者 9 。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-69-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximum69Number(num int) int {
	newNum := num
	var array []int
	for newNum > 0 {
		array = append(array, newNum%10)
		newNum /= 10
	}
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == 6 {
			val := 3
			for ; i > 0; i-- {
				val *= 10
			}
			return num + val
		}
	}
	return num
}

func Testmaximum69Number() {
	tests := []struct {
		num int
		res int
	}{
		{
			9999,
			9999,
		},
		{
			6999,
			9999,
		},
		{
			9699,
			9999,
		},
		{
			9969,
			9999,
		},
		{
			9996,
			9999,
		},
		{
			9966,
			9996,
		},
		{
			9696,
			9996,
		},
		{
			996,
			999,
		},
		{
			969,
			999,
		},
		{
			966,
			996,
		},
		{
			696,
			996,
		},
		{
			699,
			999,
		},
		{
			69,
			99,
		},
		{
			99,
			99,
		},
		{
			66,
			96,
		},
		{
			9,
			9,
		},
		{
			6,
			9,
		},
	}

	for _, test := range tests {
		res := maximum69Number(test.num)
		fmt.Println(res == test.res)
	}
}
