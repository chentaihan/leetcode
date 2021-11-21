package dynamicTag

import "fmt"

/**
1493. 删掉一个元素以后全为 1 的最长子数组
给你一个二进制数组 nums ，你需要从中删掉一个元素。

请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

如果不存在这样的子数组，请返回 0 。

 

提示 1：

输入：nums = [1,1,0,1]
输出：3
解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
示例 2：

输入：nums = [0,1,1,1,0,1,1,0,1]
输出：5
解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
示例 3：

输入：nums = [1,1,1]
输出：2
解释：你必须要删除一个元素。
示例 4：

输入：nums = [1,1,0,0,1,1,1,0,1]
输出：4
示例 5：

输入：nums = [0,0,0]
输出：0
 

提示：

1 <= nums.length <= 10^5
nums[i] 要么是 0 要么是 1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-subarray-of-1s-after-deleting-one-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestSubarray(nums []int) int {
	list := make([]int, 0, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count := 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 0 {
					count++
				} else {
					break
				}
			}
			i += count - 1
			if count > 1 || len(list) == 0 || i == len(nums)-1 {
				list = append(list, 0)
			}
		} else {
			count := 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 1 {
					count++
				} else {
					break
				}
			}
			i += count - 1
			list = append(list, count)
		}
	}
	if len(list) == 1 {
		if list[0] == 0 {
			return 0
		}
		return list[0] - 1
	}
	max := list[0] + list[1]
	for i := 1; i < len(list); i++ {
		if list[i]+list[i-1] > max {
			max = list[i] + list[i-1]
		}
	}
	return max
}

func TestlongestSubarray() {
	tests := []struct {
		nums []int
		res  int
	}{
		{
			[]int{0, 1, 0},
			1,
		},
		{
			[]int{1, 0},
			1,
		},
		{
			[]int{0, 1},
			1,
		},
		{
			[]int{1, 1, 0, 1},
			3,
		},
		{
			[]int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			5,
		},
		{
			[]int{1, 1, 1},
			2,
		},
		{
			[]int{0, 0, 0},
			0,
		},
		{
			[]int{1, 1, 0, 0, 1, 1, 1, 0, 1},
			4,
		},
	}
	for _, test := range tests {
		res := longestSubarray(test.nums)
		fmt.Println(res == test.res)
	}
}
