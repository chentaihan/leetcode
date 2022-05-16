package recursionTag

import (
	"fmt"
)

/**
1248. 统计「优美子数组」
给你一个整数数组 nums 和一个整数 k。

如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。

请返回这个数组中「优美子数组」的数目。



示例 1：

输入：nums = [1,1,2,1,1], k = 3
输出：2
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
示例 2：

输入：nums = [2,4,6], k = 1
输出：0
解释：数列中不包含任何奇数，所以不存在优美子数组。
示例 3：

输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
输出：16


提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10^5
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numberOfSubarrays(nums []int, k int) int {
	var array []int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			array = append(array, i)
		}
	}
	if len(array) < k {
		return 0
	}
	array = append(array, len(nums))
	count := 0
	start, end := 0, k-1
	for i := k; i < len(array); i++ {
		end = i
		left := array[start] + 1
		if start > 0 {
			left = array[start] - array[start-1]
		}
		right := array[end] - array[end-1]
		count += left * right
		start++
	}
	return count
}

func TestnumberOfSubarrays() {
	tests := []struct {
		nums   []int
		k      int
		result int
	}{
		{
			[]int{1, 1, 2, 1, 1},
			1,
			6,
		},
		{
			[]int{1, 1, 2, 1, 2, 1},
			1,
			9,
		},
		{
			[]int{1, 1, 2, 1, 1},
			3,
			2,
		},
		{
			[]int{2, 4, 6},
			1,
			0,
		},
		{
			[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			2,
			16,
		},
		{
			[]int{1, 1, 2, 1, 1},
			2,
			5,
		},
	}
	for _, test := range tests {
		result := numberOfSubarrays(test.nums, test.k)
		fmt.Println(test.result == result)
	}
}
