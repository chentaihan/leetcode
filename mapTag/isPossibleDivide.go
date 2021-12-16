package mapTag

import (
	"fmt"
	"sort"
)

/**
1296. 划分数组为连续数字的集合
给你一个整数数组 nums 和一个正整数 k，请你判断是否可以把这个数组划分成一些由 k 个连续数字组成的集合。
如果可以，请返回 True；否则，返回 False。

注意：此题目与 846 重复：https://leetcode-cn.com/problems/hand-of-straights/

示例 1：

输入：nums = [1,2,3,3,4,4,5,6], k = 4
输出：true
解释：数组可以分成 [1,2,3,4] 和 [3,4,5,6]。
示例 2：

输入：nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
输出：true
解释：数组可以分成 [1,2,3] , [2,3,4] , [3,4,5] 和 [9,10,11]。
示例 3：

输入：nums = [3,3,2,2,1,1], k = 3
输出：true
示例 4：

输入：nums = [1,2,3,4], k = 3
输出：false
解释：数组不能分成几个大小为 3 的子数组。


提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= nums.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-array-in-sets-of-k-consecutive-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if m[val] > 0 {
			m[val]--
			for j := 1; j < k; j++ {
				if m[val+j] <= 0 {
					return false
				}
				m[val+j]--
			}
		}
	}
	return true
}

func TestisPossibleDivide() {
	tests := []struct {
		nums []int
		k    int
		res  bool
	}{
		//{
		//	[]int{16, 21, 26, 35}, 4, false,
		//},
		{
			[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4, true,
		},
		{
			[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3, true,
		},
		{
			[]int{3, 3, 2, 2, 1, 1}, 3, true,
		},
		{
			[]int{1, 2, 3, 4}, 3, false,
		},
	}
	for _, test := range tests {
		res := isPossibleDivide(test.nums, test.k)
		fmt.Println(res == test.res)
	}
}
