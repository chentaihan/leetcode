package arrayTag

import "fmt"

/**
1551. 使数组中所有元素相等的最小操作数
存在一个长度为 n 的数组 arr ，其中 arr[i] = (2 * i) + 1 （ 0 <= i < n ）。

一次操作中，你可以选出两个下标，记作 x 和 y （ 0 <= x, y < n ）并使 arr[x] 减去 1 、arr[y] 加上 1 （即 arr[x] -=1 且 arr[y] += 1 ）。最终的目标是使数组中的所有元素都 相等 。题目测试用例将会 保证 ：在执行若干步操作后，数组中的所有元素最终可以全部相等。

给你一个整数 n，即数组的长度。请你返回使数组 arr 中所有元素相等所需的 最小操作数 。

示例 1：

输入：n = 3
输出：2
解释：arr = [1, 3, 5]
第一次操作选出 x = 2 和 y = 0，使数组变为 [2, 3, 4]
第二次操作继续选出 x = 2 和 y = 0，数组将会变成 [3, 3, 3]
示例 2：

输入：n = 6
输出：9

提示：
1 <= n <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-operations-to-make-array-equal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：
1 => 1 => 0
1, 3 => 2 >= 1			 1
1,3,5 => 3 >= 2			 2
1,3,5,7 => 4 >= 4    	 1,3
1,3,5,7,9 => 5 >= 6      2,4
1,3,5,7,9,11 => 6 >= 9   1,3,5
*/

func minOperations(n int) int {
	return n*n/4
}

func TestminOperations() {
	tests := []struct {
		n int
		r int
	}{
		{
			1, 0,
		},
		{
			2, 1,
		},
		{
			3, 2,
		},
		{
			4, 4,
		},
		{
			5, 6,
		},
		{
			6, 9,
		},
	}
	for _, test := range tests {
		r := minOperations(test.n)
		fmt.Println(r == test.r)
	}
}
