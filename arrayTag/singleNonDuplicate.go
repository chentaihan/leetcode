package arrayTag

import "fmt"

/**
540. 有序数组中的单一元素
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

 

示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10
 

提示:

1 <= nums.length <= 105
0 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-element-in-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func singleNonDuplicate(nums []int) int {
	start, end, mid := 0, len(nums)-1, 0
	for start <= end {
		mid = (start + end) / 2
		if (mid == 0 || mid == len(nums)-1) || (nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1]) {
			return nums[mid]
		} else if (mid%2 == 0 && nums[mid] != nums[mid+1]) || (mid%2 == 1 && nums[mid] != nums[mid-1]) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return nums[mid]
}

func TestsingleNonDuplicate() {
	tests := []struct {
		nums []int
		res  int
	}{
		{
			[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2,
		},
		{
			[]int{1, 1, 2, 2, 3, 4, 4, 8, 8}, 3,
		},
		{
			[]int{1, 1, 2, 2, 3, 3, 4, 8, 8}, 4,
		},
		{
			[]int{3, 3, 7, 7, 10, 11, 11}, 10,
		},
		{
			[]int{1, 1, 2}, 2,
		},
		{
			[]int{ 2}, 2,
		},
		{
			[]int{1, 2, 2}, 1,
		},
		{
			[]int{1, 1, 2, 2, 3}, 3,
		},
		{
			[]int{1, 1, 2, 3, 3}, 2,
		},
		{
			[]int{1, 2, 2, 3, 3}, 1,
		},
	}
	for _, test := range tests {
		res := singleNonDuplicate(test.nums)
		fmt.Println(res == test.res)
	}
}
