package title

import "fmt"

/**
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
 */

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	curPos, pos0, pos2 := 0, 0, len(nums)-1
	for curPos <= pos2 {
		for nums[pos2] == 2 {
			pos2--
			if pos2 == 0 || curPos > pos2 {
				return
			}
		}

		if nums[curPos] == 2 {
			nums[curPos], nums[pos2] = nums[pos2], nums[curPos]
			pos2--
		}
		if nums[curPos] == 0 {
			nums[curPos], nums[pos0] = nums[pos0], nums[curPos] //交换后nums[curPos]不可能==2|0
			pos0++
		}
		curPos++ //大胆的++
	}
}

func intArrayEqual(nums1, nums2 []int) bool {
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

func TestSortColors() {
	tests := []struct {
		sour []int
		dest []int
	}{
		{
			[]int{0, 2, 1, 2, 0, 2, 0, 1},
			[]int{0, 0, 0, 1, 1, 2, 2, 2},
		},
		{
			[]int{1, 2, 1, 2, 0, 1, 0, 0},
			[]int{0, 0, 0, 1, 1, 1, 2, 2},
		},
		{
			[]int{1, 1, 1, 1, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 1, 1, 1, 1},
		},
		{
			[]int{2, 2, 1, 1, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 2, 2, 2, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 2, 2, 2, 2},
		},
		{
			[]int{1, 1, 1, 1, 2, 2, 2, 2},
			[]int{1, 1, 1, 1, 2, 2, 2, 2},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			[]int{2, 0},
			[]int{0, 2},
		},
		{
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			[]int{2, 1, 0},
			[]int{0, 1, 2},
		},
		{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
		},
		{
			[]int{2, 2},
			[]int{2, 2},
		},
		{
			[]int{1, 1},
			[]int{1, 1},
		},
		{
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			[]int{2, 2, 0},
			[]int{0, 2, 2},
		},
		{
			[]int{1, 0, 1, 1, 1, 1, 0, 2, 2, 1, 0, 0, 0, 1, 2, 2, 1, 2, 1, 0, 0, 2, 2, 2, 0, 1, 2, 0, 1, 2, 2, 2, 2, 0, 2, 2, 2, 2, 1, 2, 1, 0, 0, 2, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
	}
	for _, test := range tests {
		sortColors(test.sour)
		fmt.Println(intArrayEqual(test.sour, test.dest))
	}
}
