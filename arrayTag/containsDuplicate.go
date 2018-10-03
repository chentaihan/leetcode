package arrayTag

/**
217. 存在重复元素
 */

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _,num := range nums {
		if _,exist := m[num]; exist {
			return true
		}else{
			m[num] = struct{}{}
		}
	}
	return false
}