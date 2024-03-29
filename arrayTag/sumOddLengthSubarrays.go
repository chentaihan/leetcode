package arrayTag

import "fmt"

/**
1588. 所有奇数长度子数组的和
给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。

子数组 定义为原数组中的一个连续子序列。

请你返回 arr 中 所有奇数长度子数组的和 。



示例 1：

输入：arr = [1,4,2,5,3]
输出：58
解释：所有奇数长度子数组和它们的和为：
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
示例 2：

输入：arr = [1,2]
输出：3
解释：总共只有 2 个长度为奇数的子数组，[1] 和 [2]。它们的和为 3 。
示例 3：

输入：arr = [10,11,12]
输出：66


提示：

1 <= arr.length <= 100
1 <= arr[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sum-of-all-odd-length-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sumOddLengthSubarrays(arr []int) int {
	count := len(arr)
	if count%2 == 0 {
		count--
	}
	total := 0
	for i := 1; i <= count; i += 2 {
		for j := 0; j < len(arr); j++ {
			start, end := j, j+i
			if end > len(arr) {
				continue
			}
			for ; start < end; start++ {
				total += arr[start]
			}
		}
	}
	return total
}

func TestsumOddLengthSubarrays() {
	tests := []struct {
		array []int
		total int
	}{
		{
			[]int{1, 4, 2, 5, 3}, 58,
		},
		{
			[]int{1, 2}, 3,
		},
		{
			[]int{10, 11, 12}, 66,
		},
	}
	for _, test := range tests {
		total := sumOddLengthSubarrays(test.array)
		if total != test.total {
			fmt.Println(total, test.total)
		} else {
			fmt.Println(true)
		}
	}
}
