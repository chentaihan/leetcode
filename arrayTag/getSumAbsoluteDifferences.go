package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
1685. 有序数组中差绝对值之和
给你一个 非递减 有序整数数组 nums 。

请你建立并返回一个整数数组 result，它跟 nums 长度相同，且result[i] 等于 nums[i] 与数组中所有其他元素差的绝对值之和。

换句话说， result[i] 等于 sum(|nums[i]-nums[j]|) ，其中 0 <= j < nums.length 且 j != i （下标从 0 开始）。



示例 1：

输入：nums = [2,3,5]
输出：[4,3,5]
解释：假设数组下标从 0 开始，那么
result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4，
result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3，
result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5。
示例 2：

输入：nums = [1,4,6,8,10]
输出：[24,15,13,15,21]


提示：

2 <= nums.length <= 105
1 <= nums[i] <= nums[i + 1] <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-absolute-differences-in-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
分析：可以分两部分计算，当前元素的结果 = 当前元素之前的结果 + 之后元素的结果

*/

func getSumAbsoluteDifferences(nums []int) []int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	total := sum[len(sum)-1]
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = (2*i+2-len(sum))*nums[i] - 2*sum[i] + total
	}
	return result
}

func TestgetSumAbsoluteDifferences() {
	tests := []struct {
		nums   []int
		result []int
	}{
		{
			[]int{1, 3},
			[]int{2, 2},
		},
		{
			[]int{1, 3, 6},
			[]int{7, 5, 8},
		},
		{
			[]int{1, 3, 5, 7, 9},
			[]int{20, 14, 12, 14, 20},
		},
		{
			[]int{2, 4, 6, 8, 10},
			[]int{20, 14, 12, 14, 20},
		},
		{
			[]int{2, 4, 6, 8},
			[]int{12, 8, 8, 12},
		},
		{
			[]int{1, 4, 6, 8},
			[]int{15, 9, 9, 13},
		},
	}
	for _, test := range tests {
		result := getSumAbsoluteDifferences(test.nums)
		fmt.Println(common.IntEqual(test.result, result))
	}
}
