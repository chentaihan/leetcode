package common

import "sort"

func IntEqual(nums1, nums2 []int) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func IntEqualSort(nums1, nums2 []int) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func IntEqualEx(nums1, nums2 [][]int) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if !IntEqual(nums1[i], nums2[i]) {
			return false
		}
	}
	return true
}

func CopyInts(array []int) []int {
	ret := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		ret[i] = array[i]
	}
	return ret
}
