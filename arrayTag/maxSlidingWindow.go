package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"sort"
)

/**
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]
示例 3：

输入：nums = [1,-1], k = 1
输出：[1,-1]
示例 4：

输入：nums = [9,11], k = 2
输出：[11]
示例 5：

输入：nums = [4,-2], k = 2
输出：[4]

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-maximum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxSlidingWindow(nums []int, k int) []int {
	index := k
	if len(nums) < index {
		index = len(nums)
	}
	type window struct {
		val   int
		index int
	}
	list := make([]window, index)
	for i := 0; i < index; i++ {
		list[i] = window{nums[i], i}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].val > list[j].val
	})
	result := []int{list[0].val}
	for ; index < len(nums); index++ {
		for len(list) > 0 && list[0].index <= index-k {
			list = list[1:]
		}
		isFind := false
		for i := len(list) - 1; i >= 0; i-- {
			if list[i].val > nums[index] {
				list = list[:i+1]
				list = append(list, window{nums[index], index})
				isFind = true
				break
			}
		}
		if !isFind {
			if len(list) == 0 {
				list = append(list, window{nums[index], index})
			} else {
				list = list[:1]
				list[0] = window{nums[index], index}
			}
		}
		result = append(result, list[0].val)
	}
	return result
}

func TestmaxSlidingWindow() {
	tests := []struct {
		nums   []int
		k      int
		result []int
	}{
		{
			[]int{9, 10, 9, -7, -4, -8, 2, -6},
			5,
			[]int{10, 10, 9, 2},
		},
		{
			[]int{1, -1},
			1,
			[]int{1, -1},
		},
		{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			[]int{4, 4, 4, 4, 4, 4, 4, 4},
			3,
			[]int{4, 4, 4, 4, 4, 4},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			[]int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			3,
			[]int{9, 8, 7, 6, 5, 4, 3},
		},
		{
			[]int{9, 8, 7, 9, 6, 5, 7, 8, 2},
			3,
			[]int{9, 9, 9, 9, 7, 8, 8},
		},
	}
	for _, test := range tests {
		result := maxSlidingWindow(test.nums, test.k)
		fmt.Println(common.IntEqual(result, test.result))
	}
}
