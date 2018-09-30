package title

import "fmt"

func removeDuplicates1(nums []int) int {
	len := len(nums)
	if len <= 1 {
		return len
	}
	curVal := nums[len-1]
	for i := len - 2; i >= 0; i-- {
		if curVal == nums[i] {
			for j := i; j < len-1; j++ {
				nums[j] = nums[j+1]
			}
			len--
		} else {
			curVal = nums[i]
		}
	}
	return len
}

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len <= 1 {
		return len
	}
	curVal := nums[0]
	savePos := 1
	for i := 1; i < len; i++ {
		if curVal != nums[i] {
			curVal = nums[i]
			nums[savePos] = nums[i]
			savePos++
		}
	}
	return savePos
}

func TestRemoveDuplicates() {
	fmt.Println(removeDuplicates([]int{}) == 0)
	fmt.Println(removeDuplicates([]int{0}) == 1)
	fmt.Println(removeDuplicates([]int{0, 0}) == 1)
	fmt.Println(removeDuplicates([]int{1, 1, 2}) == 2)
	fmt.Println(removeDuplicates([]int{1, 2, 3}) == 3)
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3}) == 3)
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3}) == 3)
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3}) == 3)
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 3}) == 3)
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3, 3}) == 3)
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4}) == 4)
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3, 3, 4, 5}) == 5)
	fmt.Println(removeDuplicates([]int{0, 1, 1, 2, 2, 3, 3, 4, 5}) == 6)
}
