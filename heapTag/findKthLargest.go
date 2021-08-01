package heapTag

import "fmt"

/**
215. 数组中的第K个最大元素
*/

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return -1
	}
	index := findKthPartion(nums)
	if index+1 == k {
		return nums[index]
	}
	if index >= k {
		return findKthLargest(nums[:index], k)
	}
	return findKthLargest(nums[index+1:], k-index-1)
}

func findKthPartion(nums []int) int {
	start, end := 0, len(nums)-1
	flag, index := nums[start], start+1
	for start < end {
		if nums[index] > flag {
			nums[index], nums[start] = nums[start], nums[index]
			start++
			index++
		} else {
			nums[index], nums[end] = nums[end], nums[index]
			end--
		}
	}
	nums[start] = flag
	return start
}

func TestfindKthLargest() {
	tests := []struct {
		nums []int
		k    int
		res  int
	}{
		{
			[]int{1, 3, 2, 9, 4, 5, 6, 8, 7, 10},
			8,
			3,
		},
		{
			[]int{1, 3, 2, 9, 4, 5, 6, 8, 7, 10},
			3,
			8,
		},
		{
			[]int{1, 3, 2, 9, 4, 5, 6, 8, 7, 10},
			4,
			7,
		},
		{
			[]int{1, 3, 2, 9, 4, 5, 6, 8, 7, 10},
			1,
			10,
		},
		{
			[]int{1, 3, 2, 9, 4, 5, 6, 8, 7, 10},
			2,
			9,
		},
		{
			[]int{1, 3, 2, 9, 4, 5, 6, 8, 7, 10},
			9,
			2,
		},
	}

	for _, test := range tests {
		res := findKthLargest(test.nums, test.k)
		fmt.Println(res == test.res, res, test.res)
	}
}
