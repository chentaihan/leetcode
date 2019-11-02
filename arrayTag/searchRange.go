package arrayTag

/**
34. 在排序数组中查找元素的第一个和最后一个位置
分析：下面的实现简单的实现时间复杂度为O(N)，可以用二分查找，排序数组查找第一时间就应该想到二分查找
由于已经打败100%golang代码，就懒得写了
*/

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	for index, val := range nums {
		if val == target {
			if result[0] == -1 {
				result[0], result[1] = index, index
			} else {
				result[1] = index
			}
		} else if result[0] != -1 {
			break
		}
	}
	return result
}
