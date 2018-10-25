package bitTag

/**
78. 子集
 */

//func subsets(nums []int) [][]int {
//	ret := [][]int{}
//	_subsets(nums, &ret)
//	return ret
//}
//
//func _subsets(nums []int, ret *[][]int){
//	if len(nums) > 0 {
//		*ret = append(*ret, nums[:1])
//		_subsets(nums[1:], ret)
//	}
//}