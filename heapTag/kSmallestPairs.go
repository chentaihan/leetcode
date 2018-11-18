package heapTag

import "fmt"

/**
373. 查找和最小的K对数字
 */


 //TODO
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	l1 := len(nums1)
	l2 := len(nums2)
	result := make([][]int, 0, k)
	if l1*l2 > k {
		index1, index2, count := 0, 0, 0
		//result = append(result, []int{nums1[0], nums2[0]})
		for index1 < l1 && index2 < l2 {
			if nums1[index1] < nums2[index2] {
				for i := 0; i <= index2 && count < k; i++ {
					result = append(result, []int{nums1[index1], nums2[i]})
					count++
				}
				index1++
			} else {
				for i := 0; i <= index1 && count < k; i++ {
					result = append(result, []int{nums1[i], nums2[index2]})
					count++
				}
				index2++
			}
			if count == k {
				break
			}
		}
	} else {
		for i := 0; i < l1; i++ {
			for j := 0; j < l2; j++ {
				result = append(result, []int{nums1[i], nums2[j]})
			}
		}
	}

	return result
}

func TestkSmallestPairs() {
	tests := []struct {
		nums1 []int
		nums2 []int
		k     int
		ret   [][]int
	}{
		{
			[]int{1, 4, 7},
			[]int{1, 2, 3},
			4,
			[][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			4,
			[][]int{{1, 1}, {1, 2}, {1, 2}, {2, 2}},
		},
	}
	for _, test := range tests {
		ret := kSmallestPairs(test.nums1, test.nums2, test.k)
		fmt.Println(ret)
	}
}
