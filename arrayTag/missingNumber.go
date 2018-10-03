package arrayTag

/**
268. 缺失数字
 */

func missingNumber(nums []int) int {
	sum := 0
	for _,val := range nums{
		sum += val
	}
	return (len(nums)+1)*len(nums)/2 -sum
}