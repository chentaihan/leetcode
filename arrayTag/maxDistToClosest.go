package arrayTag

import "fmt"

/**
849. 到最近的人的最大距离
分析：(中间连续空位-1)/2,两边就是空位数量
 */

func maxDistToClosest(seats []int) int {
	startCount, cenCount, maxCenCount := 0, 0, 0
	i := 0
	for ; i < len(seats); i++ {
		if seats[i] == 0 {
			startCount++
		} else {
			break
		}
	}
	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			cenCount = 0
		} else {
			cenCount++
			if cenCount > maxCenCount {
				maxCenCount = cenCount
			}
		}
	}
	if cenCount > startCount {
		startCount = cenCount
	}
	cenCount = (maxCenCount-1)/2 + 1
	if cenCount > startCount {
		return cenCount
	}
	return startCount
}

func TestmaxDistToClosest() {
	tests := []struct {
		seats  []int
		result int
	}{
		{
			[]int{1, 0, 0, 0, 1, 0, 1},
			2,
		},
		{
			[]int{0, 0, 0, 1, 0, 0, 0, 1, 0, 1},
			3,
		},
		{
			[]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0},
			3,
		},
	}
	for _, test := range tests {
		result := maxDistToClosest(test.seats)
		fmt.Println(result == test.result)
	}
}
