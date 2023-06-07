package mapTag

import "fmt"

/**
1496. 判断路径是否相交
给你一个字符串 path，其中 path[i] 的值可以是 'N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。

你从二维平面上的原点 (0, 0) 处开始出发，按 path 所指示的路径行走。

如果路径在任何位置上与自身相交，也就是走到之前已经走过的位置，请返回 true ；否则，返回 false 。



示例 1：



输入：path = "NES"
输出：false
解释：该路径没有在任何位置相交。
示例 2：



输入：path = "NESWW"
输出：true
解释：该路径经过原点两次。


提示：

1 <= path.length <= 104
path[i] 为 'N'、'S'、'E' 或 'W'

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-crossing
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPathCrossing(path string) bool {
	m := make(map[Point]struct{})
	point := Point{}
	m[point] = struct{}{}
	for i := 0; i < len(path); i++ {
		dt := direct(path[i])
		point.X += dt.X
		point.Y += dt.Y
		if _, exist := m[point]; exist {
			return true
		}
		m[point] = struct{}{}
	}
	return false
}

type Point struct {
	X int
	Y int
}

func direct(c byte) Point {
	switch c {
	case 'N':
		return Point{0, 1}
	case 'S':
		return Point{0, -1}
	case 'E': //东
		return Point{1, 0}
	default: //西
		return Point{-1, 0}
	}
}

func TestIsPathCrossing() {
	tests := []struct {
		path string
		res  bool
	}{
		{
			"SS",
			false,
		},
		{
			"NN",
			false,
		},
		{
			"NEWS",
			true,
		},
	}
	for _, test := range tests {
		res := isPathCrossing(test.path)
		if res == test.res {
			fmt.Println(true)
		} else {
			fmt.Println(res, test.res)
		}
	}
}
