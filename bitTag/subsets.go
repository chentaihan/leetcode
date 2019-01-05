package bitTag

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

/**
78. 子集
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。

 */

func subsets(nums []int) [][]int {
	ret := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		rLen := len(ret)
		for j := 0; j < rLen; j++ {
			array := common.CopyInts(ret[j])
			array = append(array, nums[i])
			ret = append(ret, array)
		}
	}
	return ret
}

func Testsubsets() {
	tests := []struct {
		nums []int
		ret  [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			[]int{1, 2},
			[][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			[]int{1},
			[][]int{{}, {1}},
		},
		{
			[]int{9, 0, 3, 5, 7},
			[][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}},
		},
		{
			[]int{3, 2, 4, 1},
			[][]int{{}, {3}, {2}, {2, 3}, {4}, {3, 4}, {2, 4}, {2, 3, 4}, {1}, {1, 3}, {1, 2}, {1, 2, 3}, {1, 4}, {1, 3, 4}, {1, 2, 4}, {1, 2, 3, 4}},
		},
	}

	for _, test := range tests {
		ret := subsets(test.nums)
		fmt.Println(common.IntEqualEx(ret, test.ret))
	}
}
