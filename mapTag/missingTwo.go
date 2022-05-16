package mapTag

/**
面试题 17.19. 消失的两个数字
给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

以任意顺序返回这两个数字均可。

示例 1:

输入: [1]
输出: [2,3]
示例 2:

输入: [2,3]
输出: [1,4]
提示：

nums.length <= 30000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-two-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func missingTwo(nums []int) []int {
	nums = append(nums, 0)
	nums = append(nums, 0)
	for i := 0; i < len(nums); i++ {
		for nums[i] != 0 && nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}
