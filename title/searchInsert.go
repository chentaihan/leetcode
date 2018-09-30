package title

import "fmt"

func searchInsert(nums []int, target int) int {
	start, middle, end := 0, 0, len(nums)- 1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle -1
		} else {
			start = middle + 1
		}
	}
	if middle < len(nums) && target > nums[middle] {
		return middle+1
	}
	return middle
}

func TestSearchInsert() {
	tests := []struct {
		nums   []int
		target int
		index  int
	}{
		{
			nil,
			0,
			0,
		},
		{
			[]int{},
			0,
			0,
		},
		{
			[]int{0},
			0,
			0,
		},
		{
			[]int{0, 3, 5, 7},
			3,
			1,
		},
		{
			[]int{0, 3, 5, 7},
			4,
			2,
		},
		{
			[]int{0, 3, 5, 7},
			0,
			0,
		},
		{
			[]int{0, 3, 5, 7},
			7,
			3,
		},
		{
			[]int{0, 3, 5, 7},
			9,
			4,
		},
		{
			[]int{0, 3, 5, 7, 9},
			3,
			1,
		},
		{
			[]int{0, 3, 5, 7, 9},
			4,
			2,
		},
		{
			[]int{0, 3, 5, 7, 9},
			0,
			0,
		},
		{
			[]int{0, 3, 5, 7, 9},
			9,
			4,
		},
		{
			[]int{0, 3, 5, 7, 9},
			11,
			5,
		},
	}
	for _, test := range tests {
		fmt.Println(searchInsert(test.nums, test.target) == test.index)
	}
}
