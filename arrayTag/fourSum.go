package arrayTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
	"sort"
)

/**
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：答案中不可以包含重复的四元组。

示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [], target = 0
输出：[]

提示：

0 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func fourSum(nums []int, target int) [][]int {
	var list [][4]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			list = append(list, [4]int{i, j, nums[i], nums[j]})
		}
	}
	var result [][]int
	m := make(map[int][][4]int, len(list))
	for i := 0; i < len(list); i++ {
		sum := list[i][2] + list[i][3]
		if val, ok := m[target-sum]; ok {
			for index := 0; index < len(val); index++ {
				if !fourSumHasEqual(list[i], val[index]) {
					item := []int{list[i][2], list[i][3], val[index][2], val[index][3]}
					sort.Ints(item)
					if !fourSumExist(result, item) {
						result = append(result, item)
					}
				}
			}
		}
		m[sum] = append(m[sum], list[i])
	}

	return result
}

func fourSumHasEqual(a, b [4]int) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if a[i] == b[j] {
				return true
			}
		}
	}
	return false
}

func fourSumExist(nums [][]int, item []int) bool {
	for i := 0; i < len(nums); i++ {
		equal := true
		for j := 0; j < 4; j++ {
			if nums[i][j] != item[j] {
				equal = false
				break
			}
		}
		if equal {
			return equal
		}
	}
	return false
}

func TestfourSum() {
	tests := []struct {
		nums   []int
		target int
		res    [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
	}

	for _, test := range tests {
		res := fourSum(test.nums, test.target)
		fmt.Println(res)
		fmt.Println(common.IntEqualEx(res, test.res))
	}
}
