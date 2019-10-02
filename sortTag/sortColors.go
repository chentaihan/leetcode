package sortTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
// 注意:
//不能使用代码库中的排序函数来解决这道题。
//
// 示例:
//
// 输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//
// 进阶：
//
//
// 一个直观的解决方案是使用计数排序的两趟扫描算法。
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
//

func sortColors(nums []int) {
	start, end, index := 0, len(nums)-1, 0
	for start < end {
		for start < end && nums[start] == 0 {
			start++
		}
		for start < end && nums[end] == 2 {
			end--
		}
		if index < start {
			index = start
		}
		if index > end {
			break
		}
		needAdd := true
		if nums[index] == 0 {
			nums[start], nums[index] = nums[index], nums[start]
			start++
			needAdd = false
		}
		if nums[index] == 2 {
			nums[end], nums[index] = nums[index], nums[end]
			end--
			needAdd = false
		}
		if needAdd {
			index++
		}
	}
}

func TestsortColors() {
	tests := []struct {
		nums  []int
		resst []int
	}{
		{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			[]int{0, 1, 0, 2},
			[]int{0, 0, 1, 2},
		},
		{
			[]int{0, 1, 0, 1, 2},
			[]int{0, 0, 1, 1, 2},
		},
		{
			[]int{2, 0, 1, 0, 1, 2},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 2, 1, 1, 0, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{0, 1, 2, 2, 2, 1, 1, 0, 0},
			[]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			[]int{0, 1, 2, 2, 2, 1, 1, 0, 0, 1},
			[]int{0, 0, 0, 1, 1, 1, 1, 2, 2, 2},
		},
		{
			[]int{1, 2, 0, 1, 2, 1, 0, 1, 2, 0, 0},
			[]int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2},
		},
		{
			[]int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2},
			[]int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2},
		},
		{
			[]int{2, 1, 0, 0, 1, 0, 1, 1, 2, 2, 0},
			[]int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2},
		},
	}

	for _, test := range tests {
		sortColors(test.nums)
		if !common.IntEqual(test.nums, test.resst) {
			fmt.Println(test.nums)
			fmt.Println(test.resst)
			fmt.Println()
		} else {
			fmt.Println("ok")
		}
	}
}
