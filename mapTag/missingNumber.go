package mapTag

import "fmt"

/**
剑指 Offer 53 - II. 0～n-1中缺失的数字
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。



示例 1:

输入: [0,1,3]
输出: 2
示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8


限制：

1 <= 数组长度 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func missingNumber(nums []int) int {
	start, mid, end := 0, 0, len(nums)-1
	for start <= end {
		mid = (start + end) / 2
		if mid == nums[mid] {
			start = mid + 1
		} else if mid < nums[mid] {
			if mid > 0 && nums[mid] > nums[mid-1]+1 {
				return nums[mid] - 1
			} else {
				end = mid - 1
			}
		}
	}
	if mid == nums[mid] {
		return mid + 1
	}
	return mid
}

func TestmissingNumber() {
	tests := []struct {
		nums []int
		res  int
	}{
		{
			[]int{0, 1, 3},
			2,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			8,
		},
		{
			[]int{0, 1, 2},
			3,
		},
		{
			[]int{1, 2},
			0,
		},
		{
			[]int{0, 1, 2, 3, 5},
			4,
		},
		{
			[]int{0, 1, 3, 4, 5},
			2,
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			6,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 7},
			6,
		},
	}

	for _, test := range tests {
		res := missingNumber(test.nums)
		fmt.Println(res == test.res)
	}
}
