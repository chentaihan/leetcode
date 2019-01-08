package mapTag

/**
349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
 */

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = false
	}

	ret := []int{}
	for i := 0; i < len(nums2); i++ {
		if val, exist := m[nums2[i]]; exist && !val {
			ret = append(ret, nums2[i])
			m[nums2[i]] = true
		}
	}
	return ret
}
