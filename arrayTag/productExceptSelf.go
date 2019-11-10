package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
238. 除自身以外数组的乘积
给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:

输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/product-of-array-except-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left := 1
	for i := 0; i < len(nums); i++ { //左边的乘积
		result[i] = left
		left *= nums[i]
	}
	right := 1
	for j := len(nums) - 1; j >= 0; j-- { //右边的乘积
		result[j] *= right
		right *= nums[j]
	}
	return result
}

func TestproductExceptSelf() {
	tests := []struct {
		nums []int
		ret  []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{120, 60, 40, 30, 24},
		},
		{
			[]int{1, 2, 3, 4, 5, 5},
			[]int{600, 300, 200, 150, 120, 120},
		},
	}
	for _, test := range tests {
		ret := productExceptSelf(test.nums)
		fmt.Println(common.IntEqual(ret, test.ret))
	}
}
