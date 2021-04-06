package depthFirstTag

/**
491. 递增子序列
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:

给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findSubsequences(nums []int) [][]int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		result = _findSubsequences(result, nums[i])
	}
	newResult := make([][]int, 0, len(result))
	for i := 0; i < len(result); i++ {
		if len(result[i]) > 1 {
			newResult = append(newResult, result[i])
		}
	}
	return newResult
}

func _findSubsequences(result [][]int, num int) [][]int {
	l := len(result)
	for i := 0; i < l; i++ {
		array := result[i]
		if array[len(array)-1] <= num {
			newArray := make([]int, len(array)+1)
			copy(newArray, array)
			newArray[len(array)] = num
			if !checkFindSubsequences(result, newArray) {
				result = append(result, newArray)
			}
		}
	}
	result = append(result, []int{num})
	return result
}

func checkFindSubsequences(result [][]int, array []int) bool {
	for i := 0; i < len(result); i++ {
		res := result[i]
		if len(res) != len(array) {
			continue
		}
		equal := true
		for j := 0; j < len(res); j++ {
			if res[j] != array[j] {
				equal = false
				break
			}
		}
		if equal {
			return true
		}
	}
	return false
}
