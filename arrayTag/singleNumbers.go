package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
剑指 Offer 56 - I. 数组中数字出现的次数
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

示例 1：

输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
示例 2：

输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func singleNumbers(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum ^= nums[i]
	}

	result := 1
	for sum&1 != 1 {
		result <<= 1
		sum >>= 1
	}

	first, second := 0, 0
	for i := 0; i < len(nums); i++ {
		if result&nums[i] > 0 {
			first ^= nums[i]
		} else {
			second ^= nums[i]
		}
	}
	return []int{first, second}
}

func TestsingleNumbers() {
	tests := []struct {
		array  []int
		result []int
	}{
		{
			[]int{2, 5, 2, 4, 3, 5, 6, 4},
			[]int{3, 6},
		},
		{
			[]int{4, 1, 6, 4},
			[]int{1, 6},
		},
		{
			[]int{2, 2, 4, 1, 6, 4},
			[]int{1, 6},
		},
	}

	for _, test := range tests {
		res := singleNumbers(test.array)
		fmt.Println(common.IntEqualSort(res, test.result))
	}
}
