package depthFirstTag

import "fmt"

/**
815. 公交路线
给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。

例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。

求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。



示例 1：

输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
输出：2
解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
示例 2：

输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
输出：-1


提示：

1 <= routes.length <= 500.
1 <= routes[i].length <= 105
routes[i] 中的所有值 互不相同
sum(routes[i].length) <= 105
0 <= routes[i][j] < 106
0 <= source, target < 106

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bus-routes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numBusesToDestination(routes [][]int, source int, target int) int {
	mList := make([]map[int]bool, len(routes))
	for i := 0; i < len(routes); i++ {
		m := make(map[int]bool, len(routes[i]))
		route := routes[i]
		for j := 0; j < len(route); j++ {
			m[route[j]] = true
		}
		mList[i] = m
	}
	dp := make([]bool, len(routes))
	minCount := -1
	for i := 0; i < len(mList); i++ {
		if mList[i][source] {
			if target == source {
				return 0
			}
			_numBusesToDestination(i, mList, dp, target, 1, &minCount)
		}
	}
	return minCount
}

func _numBusesToDestination(index int, mList []map[int]bool, dp []bool, target, count int, minCount *int) {
	route := mList[index]
	if route[target] {
		if *minCount == -1 {
			*minCount = count
		} else if *minCount > count {
			*minCount = count
		}
		return
	}
	if dp[index] {
		return
	}
	dp[index] = true
	for i := 0; i < len(dp); i++ {
		if dp[i] {
			continue
		}
		for key := range route {
			if mList[i][key] {
				_numBusesToDestination(i, mList, dp, target, count+1, minCount)
			}
		}
	}
	dp[index] = false
}

func TestnumBusesToDestination() {
	tests := []struct {
		routes [][]int
		source int
		target int
		result int
	}{
		{
			[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6, 2,
		},
		{
			[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 7, 1,
		},
		{
			[][]int{{1, 2, 7}, {3, 6, 7}}, 3, 7, 1,
		},
		{
			[][]int{{1, 2, 7}, {3, 6, 7}}, 3, 8, -1,
		},
		{
			[][]int{{1, 3, 7}, {3, 6, 8}, {6, 8, 9}}, 1, 9, 3,
		},
		{
			[][]int{{1, 3, 6}, {2, 6, 10}, {8, 10, 12}}, 1, 12, 3,
		},
		{
			[][]int{{1, 3, 6}, {2, 6, 12}, {2, 8, 12}}, 1, 12, 2,
		},
		{
			[][]int{{1, 3, 6, 12}, {2, 6, 12}, {2, 8, 12}}, 1, 12, 1,
		},
		{
			[][]int{{1, 3, 6}, {6, 7, 8}, {8, 9, 10}, {3, 10}}, 1, 10, 2,
		},
	}
	for _, test := range tests {
		result := numBusesToDestination(test.routes, test.source, test.target)
		if result != test.result {
			fmt.Println(false, result, test.result)
		} else {
			fmt.Println(result == test.result)
		}

	}
}
