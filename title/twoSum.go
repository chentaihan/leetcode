package title

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	m := make(map[int]int, len(nums))
	for index, num := range nums {
		m[num] = index
	}
	for index1, num := range nums {
		if index2, exist := m[target-num]; exist && index1 != index2 {
			return []int{index1, index2}
		}
	}
	return nil
}

func TwoSumTest() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 15
	result := TwoSum(nums, target)
	if len(result) == 2 {
		fmt.Println(result[0], result[1])
	}
	nums = []int{3, 2, 4}
	target = 6
	result = TwoSum(nums, target)
	if len(result) == 2 {
		fmt.Println(result[0], result[1])
	}
}
