package arrayTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
15. 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
 */

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	l := len(nums)
	m := make(map[int]int, len(nums))
	hashMap := map[int]struct{}{}
	result := [][]int{}
	for i := 0; i < l; i++ {
		m[nums[i]]++
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			val := - nums[i] - nums[j]
			if count, exist := m[val]; exist {
				if count == 1 && (val == nums[i] || val == nums[j]) {
					continue
				}else if count == 2 && (val == nums[i] && val == nums[j]) {
					continue
				}
				array := threeSort(nums[i], nums[j], val)
				hashCode := getHashCode(array)
				if _, exist := hashMap[hashCode]; !exist {
					result = append(result, array)
					hashMap[hashCode] = struct{}{}
				}
			}
		}
	}
	return result
}

func getHashCode(array []int) int {
	code := 5381
	for i := 0; i < len(array); i++ {
		code = code << 5 + code + array[i]
	}
	return code
}

func threeSort(num1, num2, num3 int) []int {
	if num1 > num2 {
		if num2 > num3 {
			return []int{num1, num2, num3}
		} else if num1 > num3 {
			return []int{num1, num3, num2}
		} else {
			return []int{num3, num1, num2}
		}
	} else {
		if num1 > num3 {
			return []int{num2, num1, num3}
		} else if num2 > num3 {
			return []int{num2, num3, num1}
		} else {
			return []int{num3, num2, num1}
		}
	}
}

func TestthreeSum() {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{-1,0,1,0},
			[][]int{{1, 0, -1}},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{1, 0, -1}, {2, -1, -1}},
		},
		{
			[]int{-1, 0, 1, 2, 2, -1, -4},
			[][]int{{1, 0, -1}, {2, -1, -1}, {2, 2, -4}},
		},
		{
			[]int{-3, -2, -1, 0, 1, 2, 3},
			[][]int{{3, 0, -3}, {2, 1, -3}, {3, -1, -2}, {2, 0, -2}, {1, 0, -1}},
		},
	}
	for _, test := range tests {
		result := threeSum(test.nums)
		fmt.Println(common.IntEqualEx(result, test.result))
	}
}
