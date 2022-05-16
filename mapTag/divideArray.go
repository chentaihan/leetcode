package mapTag

/**
2206. 将数组划分成相等数对
给你一个整数数组 nums ，它包含 2 * n 个整数。

你需要将 nums 划分成 n 个数对，满足：

每个元素 只属于一个 数对。
同一数对中的元素 相等 。
如果可以将 nums 划分成 n 个数对，请你返回 true ，否则返回 false 。



示例 1：

输入：nums = [3,2,3,2,2,2]
输出：true
解释：
nums 中总共有 6 个元素，所以它们应该被划分成 6 / 2 = 3 个数对。
nums 可以划分成 (2, 2) ，(3, 3) 和 (2, 2) ，满足所有要求。
示例 2：

输入：nums = [1,2,3,4]
输出：false
解释：
无法将 nums 划分成 4 / 2 = 2 个数对且满足所有要求。


提示：

nums.length == 2 * n
1 <= n <= 500
1 <= nums[i] <= 500

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/divide-array-into-equal-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func divideArray(nums []int) bool {
	if len(nums)&1 == 1 {
		return false
	}
	m := make(map[int]int, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, count := range m {
		if count&1 == 1 {
			return false
		}
	}
	return true
}
