package stringTag

import (
	"fmt"
	"sort"
)

/**
539. 最小时间差
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
示例 1：

输入：timePoints = ["23:59","00:00"]
输出：1
示例 2：

输入：timePoints = ["00:00","23:59","00:00"]
输出：0


提示：

2 <= timePoints <= 2 * 104
timePoints[i] 格式为 "HH:MM"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-time-difference
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	min := 24 * 60
	for i := 1; i < len(timePoints); i++ {
		diff := timeMinus(timePoints[i-1], timePoints[i])
		if diff < min {
			min = diff
		}
	}
	diff := timeMinus(timePoints[len(timePoints)-1], "24:00")
	if diff < min {
		time1 := timePoints[0]
		diff += int(time1[0]-'0')*600 + int(time1[1]-'0')*60 + int(time1[3]-'0')*10 + int(time1[4]-'0')
		if diff < min {
			return diff
		}
	}
	return min
}

func timeMinus(time1, time2 string) int {
	hour := int(time2[0]-'0')*10 + int(time2[1]-'0') - int(time1[0]-'0')*10 - int(time1[1]-'0')
	second := int(time2[3]-'0')*10 + int(time2[4]-'0') - int(time1[3]-'0')*10 - int(time1[4]-'0')
	return hour*60 + second
}

func TestfindMinDifference() {
	tests := []struct {
		array  []string
		result int
	}{
		{
			[]string{"23:59", "00:00"},
			1,
		},
		{
			[]string{"23:59", "01:00", "01:01"},
			1,
		},
		{
			[]string{"00:00", "23:59", "00:00"},
			0,
		},
		{
			[]string{"00:00", "10:00"},
			600,
		},
		{
			[]string{"10:00", "11:00"},
			60,
		},
		{
			[]string{"10:00", "11:00", "11:20"},
			20,
		},
		{
			[]string{"10:00", "11:00", "11:50", "12:00"},
			10,
		},
	}

	for _, test := range tests {
		result := findMinDifference(test.array)
		fmt.Println(result == test.result)
	}

}
