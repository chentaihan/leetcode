package stackTag

import "fmt"

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//topK
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
//

func topK(nums []int, k int) int {
	index := findTopkEx(nums, 0, len(nums)-1)
	if index+1 == k {
		return nums[index]
	}
	if index >= k { //在top index个数中查找第k大的数
		return topK(nums[0:index], k)
	} else { //由于top index的数量小于k，需要在剩余的数中找到第k-index-1大的数
		return topK(nums[index+1:], k-index-1)
	}
}

//取第一个数作为分个数，它的左边数都大于等于它，它右边的数都小于它，并返回对应的索引
func findTopk(nums []int, start, end int) int {
	index := start
	flag := nums[index]
	for {
		for end > index && nums[end] < flag {
			end--
		}
		for index <= end && nums[index] >= flag {
			index++
		}
		if index < end {
			nums[index], nums[end] = nums[end], nums[index]
			index++
			end--
		} else {
			if index > 0 {
				nums[index-1], nums[start] = nums[start], nums[index-1]
			}
			return index - 1
		}
	}
	return index
}

func findTopkEx(nums []int, start, end int) int {
	index := start
	key := nums[start]
	for {
		for index < end && nums[end] < key {
			end--
		}
		for index <= end && nums[index] >= key {
			index++
		}

		if index < end {
			nums[index], nums[end] = nums[end], nums[index]
			index++
			end--
		} else {
			index--
			nums[index], nums[start] = nums[start], nums[index]
			return index
		}
	}
	return index
}

func TestfindTopk() {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			[]int{4, 3, 5, 7, 9, 2, 1, 6, 8},
			5,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			8,
		},
		{
			[]int{9, 3, 5, 7, 1, 2, 4, 6, 8},
			0,
		},
		{
			[]int{6, 3, 5, 7, 9, 2, 4, 1, 8},
			3,
		},
		{
			[]int{2, 3, 5, 7, 9, 1, 4, 6, 8},
			7,
		},
		{
			[]int{8, 3, 5, 7, 9, 2, 4, 6, 1},
			1,
		},
	}

	for _, test := range tests {
		result := findTopk(test.nums, 0, len(test.nums)-1)
		fmt.Println(result == test.result)
	}
}

func TestTopk() {
	tests := []struct {
		nums   []int
		k      int
		result int
	}{
		{
			[]int{2, 1},
			1,
			2,
		},
		{
			[]int{2, 1},
			2,
			1,
		},
		{
			[]int{1},
			1,
			1,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			1,
			9,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			2,
			8,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			3,
			7,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			4,
			6,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			5,
			5,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			6,
			4,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			7,
			3,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			8,
			2,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			9,
			1,
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			6,
			4,
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			9,
			1,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			1,
		}, {
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			8,
			2,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			9,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
			8,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 8, 8, 9},
			2,
			8,
		},
		{
			[]int{8, 8, 8, 8, 8, 8, 8, 8, 8},
			2,
			8,
		},
		{
			[]int{8, 8, 8, 8, 8, 8, 8, 8, 8},
			1,
			8,
		},
		{
			[]int{8, 8, 8, 8, 8, 8, 8, 8, 8},
			9,
			8,
		},
	}

	for _, test := range tests {
		result := topK(test.nums, test.k)
		fmt.Println(result == test.result)
	}
}

func findKthLargest(nums []int, k int) int {
	buffer := make([]int, k)
	count := 0
	i := 0
	for ; i < k; i++ {
		j := 0
		for ; j < count; j++ {
			if nums[i] < buffer[j] {
				for index := count - 1; index >= j; index-- {
					buffer[index+1] = buffer[index]
				}
				break
			}
		}
		buffer[j] = nums[i]
		count++
	}

	for ; i < len(nums); i++ {
		if nums[i] <= buffer[0] {
			continue
		}
		j := 0
		for ; j < count; j++ {
			if nums[i] > buffer[j] {
				continue
			}
			break
		}
		j--
		for index := 0; index < j; index++ {
			buffer[index] = buffer[index+1]
		}
		buffer[j] = nums[i]
	}
	return buffer[0]
}

func TestfindKthLargest() {
	tests := []struct {
		nums   []int
		k      int
		result int
	}{
		{
			[]int{2, 1},
			1,
			2,
		},
		{
			[]int{2, 1},
			2,
			1,
		},
		{
			[]int{1},
			1,
			1,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			1,
			9,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			2,
			8,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			3,
			7,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			4,
			6,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			5,
			5,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			6,
			4,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			7,
			3,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			8,
			2,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8},
			9,
			1,
		},
	}

	for _, test := range tests {
		result := findKthLargest(test.nums, test.k)
		fmt.Println(result == test.result)
	}
}
