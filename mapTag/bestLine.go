package mapTag

import "fmt"

/**
面试题 16.14. 最佳直线 TODO
给定一个二维平面及平面上的 N 个点列表Points，其中第i个点的坐标为Points[i]=[Xi,Yi]。请找出一条直线，其通过的点的数目最多。

设穿过最多点的直线所穿过的全部点编号从小到大排序的列表为S，你仅需返回[S[0],S[1]]作为答案，若有多条直线穿过了相同数量的点，则选择S[0]值较小的直线返回，S[0]相同则选择S[1]值较小的直线返回。

示例：

输入： [[0,0],[1,1],[1,0],[2,0]]
输出： [0,2]
解释： 所求直线穿过的3个点的编号为[0,2,3]
提示：

2 <= len(Points) <= 300
len(Points[i]) = 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-line-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func bestLine(points [][]int) []int {
	m := make(map[float64]*[3]int, len(points))
	var max = [3]int{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x := points[j][0] - points[i][0]
			y := points[j][1] - points[i][1]
			var key float64
			if x != 0 {
				key = float64(y) / float64(x)
			}
			if val, exist := m[key]; exist {
				(*val)[2]++
				if (*val)[2] > max[2] {
					max = *val
				}
			} else {
				m[key] = &[3]int{i, j, 1}
				max = [3]int{i, j, 1}
			}
		}
	}

	return max[:2]
}

func TestbestLine() {
	tests := []struct {
		points [][]int
		result []int
	}{
		{
			[][]int{{-24272, -29606}, {-37644, -4251}, {2691, -22513}, {-14592, -33765}, {-21858, 28550}, {-22264, 41303}, {-6960, 12785}, {-39133, -41833}, {25151, -26643}, {-19416, 28550}, {-17420, 22270}, {-8793, 16457}, {-4303, -25680}, {-14405, 26607}, {-49083, -26336}, {22629, 20544}, {-23939, -25038}, {-40441, -26962}, {-29484, -30503}, {-32927, -18287}, {-13312, -22513}, {15026, 12965}, {-16361, -23282}, {7296, -15750}, {-11690, -21723}, {-34850, -25928}, {-14933, -16169}, {23459, -9358}, {-45719, -13202}, {-26868, 28550}, {4627, 16457}, {-7296, -27760}, {-32230, 8174}, {-28233, -8627}, {-26520, 28550}, {5515, -26001}, {-16766, 28550}, {21888, -3740}, {1251, 28550}, {15333, -26322}, {-27677, -19790}, {20311, 7075}, {-10751, 16457}, {-47762, -44638}, {20991, 24942}, {-19056, -11105}, {-26639, 28550}, {-19862, 16457}, {-27506, -4251}, {-20172, -5440}, {-33757, -24717}, {-9411, -17379}, {12493, 29906}, {0, -21755}, {-36885, -16192}, {-38195, -40088}, {-40079, 7667}, {-29294, -34032}, {-55968, 23947}, {-22724, -22513}, {20362, -11530}, {-11817, -23957}, {-33742, 5259}, {-10350, -4251}, {-11690, -22513}, {-20241, -22513}},
			[]int{4, 9},
		},
	}
	for _, test := range tests {
		res := bestLine(test.points)
		fmt.Println(res)
	}
}
