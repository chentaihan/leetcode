package arrayTag

/**
728. 自除数
 */

func selfDividingNumbers(left int, right int) []int {
	buffer := make([]int, 0, right-left)
	for ; left <= right; left++ {
		curNum := left
		for {
			val := curNum % 10
			if val == 0 {
				goto end
			}
			if left%val != 0 {
				goto end
			}
			curNum = curNum / 10
		}
		buffer = append(buffer, left)
		end :
	}
	return buffer
}
