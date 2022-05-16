package backtrackTag

import (
	"fmt"
	"github.com/chentaihan/leetcode/common"
)

/**
797. 所有可能的路径
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）

二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。

译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。



示例 1：



输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
示例 2：



输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
示例 3：

输入：graph = [[1],[]]
输出：[[0,1]]
示例 4：

输入：graph = [[1,2,3],[2],[3],[]]
输出：[[0,1,2,3],[0,2,3],[0,3]]
示例 5：

输入：graph = [[1,3],[2],[3],[]]
输出：[[0,1,2,3],[0,3]]


提示：

n == graph.length
2 <= n <= 15
0 <= graph[i][j] < n
graph[i][j] != i（即，不存在自环）
graph[i] 中的所有元素 互不相同
保证输入为 有向无环图（DAG）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-paths-from-source-to-target
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0, len(graph))
	list := make([]int, 1, len(graph))
	_allPathsSourceTarget(graph, 0, &result, list)
	return result
}

func _allPathsSourceTarget(graph [][]int, start int, result *[][]int, list []int) {
	if list[len(list)-1] == len(graph)-1 {
		newList := make([]int, len(list))
		copy(newList, list)
		*result = append(*result, newList)
		return
	}
	row := graph[start]
	for i := 0; i < len(row); i++ {
		list = append(list, row[i])
		_allPathsSourceTarget(graph, row[i], result, list)
		list = list[:len(list)-1]
	}
}

func TestAllPathsSourceTarget() {
	tests := []struct {
		graph [][]int
		res   [][]int
	}{
		{
			[][]int{{1, 2}, {3}, {3}, {}},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		}, {
			[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		}, {
			[][]int{{1}, {}},
			[][]int{{0, 1}},
		}, {
			[][]int{{1, 2, 3}, {2}, {3}, {}},
			[][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}},
		}, {
			[][]int{{1, 3}, {2}, {3}, {}},
			[][]int{{0, 1, 2, 3}, {0, 3}},
		},
	}
	for _, test := range tests {
		res := allPathsSourceTarget(test.graph)
		fmt.Println(common.IntEqualEx(res, test.res))
	}
}
