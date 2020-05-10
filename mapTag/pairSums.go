package mapTag

/**
面试题 16.24. 数对和
设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。

示例 1:

输入: nums = [5,6,5], target = 11
输出: [[5,6]]
示例 2:

输入: nums = [5,6,5,6], target = 11
输出: [[5,6],[5,6]]
提示：

nums.length <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pairs-with-sum-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func pairSums(nums []int, target int) [][]int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var result [][]int
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 && m[target-nums[i]] > 0 {
			m[nums[i]]--
			m[target-nums[i]]--
			if m[nums[i]] >= 0 && m[target-nums[i]] >= 0 {
				result = append(result, []int{nums[i], target - nums[i]})
			}
		}
	}
	return result
}
