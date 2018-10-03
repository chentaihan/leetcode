package arrayTag

/**
896. 单调数列
 */

func isMonotonic(A []int) bool {
	i := 1
	l := len(A)
	for i < l && A[i-1] == A[i] {
		i++
	}
	if i >= l {
		return true
	}
	if A[i-1] < A[i] {
		for i++; i < l; i++ {
			if A[i-1] > A[i] {
				return false
			}
		}
	} else {
		for i++; i < l; i++ {
			if A[i-1] < A[i] {
				return false
			}
		}
	}
	return true
}
