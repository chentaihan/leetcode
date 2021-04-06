package arrayTag

import "fmt"

func TestIsSortArray() {
	tests := []struct {
		array  []int
		result bool
	}{
		{
			[]int{0, 1, 3, 4, 2, 5, 6},
			true,
		},
		{
			[]int{0, 1, 3, 4, 2, 5, 2},
			false,
		},
		{
			[]int{0, 1, 3, 4, 2, 5, 6},
			true,
		},
		{
			[]int{0, 1, 6, 4, 2, 5, 3},
			true,
		},
		{
			[]int{6, 4, 2, 5, 3, 1, 0},
			true,
		},
		{
			[]int{0, 0, 0, 0, 0, 0},
			false,
		},
		{
			[]int{1, 1, 1, 1, 1, 1},
			false,
		},
	}

	for index, test := range tests {
		fmt.Println(index, test.result == IsSortArray(test.array))
	}
}

//0,1,3,4,2,5,6
func IsSortArray(array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] != i && array[i] == array[array[i]] {
			return false
		}
		for array[i] != array[array[i]] {
			array[i], array[array[i]] = array[array[i]], array[i]
		}
	}
	return true
}

