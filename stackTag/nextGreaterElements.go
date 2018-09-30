package stackTag

/**
503. 下一个更大元素 II
 */

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	m := make(map[int]int, len(nums))
	for key, val := range nums {
		m[val] = key
	}
	buffer := make([]int, len(nums))
	for key, val := range nums {
		nextVal := -1
		for start := m[val]; start < len(nums); start++ {
			if nums[start] > val {
				nextVal = nums[start]
				break
			}
		}
		buffer[key] = nextVal
	}
	return buffer
}