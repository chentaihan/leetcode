package arrayTag

import "fmt"

/**
153. 寻找旋转排序数组中的最小值(done)
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMin(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	mid := 0
	for start <= end {
		mid = (start + end) >> 1
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return nums[mid]
}

func TestFindMin() {
	tests := []struct {
		nums []int
		min  int
	}{
		{
			[]int{2, 3, 4, 5, 6, 7, 8, 1},
			1,
		},
		{
			[]int{3, 4, 5, 6, 7, 8, 1, 2},
			1,
		},
		{
			[]int{4, 5, 6, 7, 8, 1, 2, 3},
			1,
		},
		{
			[]int{5, 6, 7, 8, 1, 2, 3, 4},
			1,
		},
		{
			[]int{6, 7, 8, 1, 2, 3, 4, 5},
			1,
		},
		{
			[]int{7, 8, 1, 2, 3, 4, 5, 6},
			1,
		},
		{
			[]int{8, 1, 2, 3, 4, 5, 6, 7},
			1,
		},
	}
	for _, test := range tests {
		min := findMin(test.nums)
		fmt.Println(min == test.min)
	}
}
