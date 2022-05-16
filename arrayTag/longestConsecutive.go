package arrayTag

import "fmt"

/**
128. 最长连续序列
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：将每个数存入map中，对应的值的左右值是否在map中，存在就将将他们连在一起，就更新连续的数量
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := map[int]int{}
	max := 1
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 {
			continue
		}
		left := m[nums[i]-1]
		count := left + 1
		right := m[nums[i]+1]
		count += right
		if left > 0 {
			m[nums[i]-left] = count
		}
		if right > 0 {
			m[nums[i]+right] = count
		}
		m[nums[i]] = count
		if count > max {
			max = count
		}
	}
	return max
}

func TestlongestConsecutive() {
	tests := []struct {
		nums []int
		val  int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
	}
	for _, test := range tests {
		val := longestConsecutive(test.nums)
		fmt.Println(val == test.val)
	}
}
