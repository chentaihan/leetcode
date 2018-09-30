package arrayTag

import "fmt"

/**
219. 存在重复元素 II(出题人是个傻逼)
 */

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 || k <= 0 {
		return false
	}
	m := map[int]int{}
	for index, num := range nums {
		if val,isOk := m[num];isOk {
			if index - val <= k {
				return true
			}
		}
		m[num] = index
	}
	return false
}

func TestcontainsNearbyDuplicate() {
	tests := []struct{
		nums []int
		k int
		result bool
	}{
		{
			[]int{1,2,3,1},
			3,
			true,
		},
		{
			[]int{1,2,1,3,1},
			3,
			true,
		},
		{
			[]int{1,2,1,3,1,1},
			5,
			true,
		},
		{
			[]int{1,0,1,1},
			3,
			true,
		},
		{
			[]int{1,2,3,1,2,3},
			2,
			false,
		},
	}
	for _,test := range tests{
		fmt.Println(containsNearbyDuplicate(test.nums, test.k) == test.result)
	}
}
