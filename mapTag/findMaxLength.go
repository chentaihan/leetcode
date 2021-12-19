package mapTag

import "fmt"

/**
525. 连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2:

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contiguous-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMaxLength(nums []int) int {
	counts := make([][2]int, len(nums))
	counts[0][nums[0]]++
	for i := 1; i < len(nums); i++ {
		counts[i][nums[i]] = counts[i-1][nums[i]] + 1
		another := (nums[i] + 1) % 2
		counts[i][another] = counts[i-1][another]
	}
	max := 0
	for i := len(counts) - 1; i >= 0; i-- {
		if counts[i][0] == counts[i][1] && i >= max {
			max = i + 1
		}
		if i <= max {
			break
		}
		for j := 0; j < i; j++ {
			count := counts[i][0] - counts[j][0]
			if count == counts[i][1]-counts[j][1] && count*2 > max {
				max = count * 2
			}
		}
	}
	return max
}

func findMaxLengthEx(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	sum, res := 0, 0
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 && i+1 > res {
			res = i + 1
		}
		if index, exist := m[sum]; exist {
			if i-index > res {
				res = i - index
			}
		} else {
			m[sum] = i
		}
	}
	return res
}

func TestfindMaxLength() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1},
			68,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{0, 1, 0, 1},
			4,
		},
		{
			[]int{0, 1, 0},
			2,
		},
		{
			[]int{1, 1, 0, 0},
			4,
		},
		{
			[]int{1, 1, 1, 0, 0},
			4,
		},
		{
			[]int{0, 1, 1, 1, 0, 0},
			6,
		},
		{
			[]int{1, 1, 0, 1, 1, 1, 0, 0},
			6,
		},
	}
	for _, test := range tests {
		result := findMaxLengthEx(test.nums)
		fmt.Println(result == test.result)
	}
}
