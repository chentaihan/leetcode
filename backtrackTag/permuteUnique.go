package backtrackTag

import (
	"fmt"
)

/**
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

 

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 

提示：

1 <= nums.length <= 8
-10 <= nums[i] <= 10


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	_permuteUnique(nums, 0, &result)
	return result
}

func _permuteUnique(nums []int, left int, result *[][]int) {
	if left == len(nums) {
		res := make([]int, len(nums))
		copy(res, nums)
		*result = append(*result, res)
	} else {
		m := make(map[int]bool, len(nums))
		for i := left; i < len(nums); i++ {
			if !m[nums[i]] {
				m[nums[i]] = true
				nums[i], nums[left] = nums[left], nums[i]
				_permuteUnique(nums, left+1, result)
				nums[i], nums[left] = nums[left], nums[i]
			}
		}
	}
}

func TestpermuteUnique() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
	nums = []int{1, 2, 3}
	fmt.Println(permuteUnique(nums))
	nums = []int{2, 2, 1, 1}
	fmt.Println(permuteUnique(nums))
	nums = []int{0, 1, 0, 0, 9}
	fmt.Println(permuteUnique(nums))
	nums = []int{-1, 2, 0, -1, 1, 0, 1}
	fmt.Println(permuteUnique(nums))

}
