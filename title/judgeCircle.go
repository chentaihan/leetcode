package title

import "fmt"

/**
初始位置 (0, 0) 处有一个机器人。给出它的一系列动作，判断这个机器人的移动路线是否形成一个圆圈，换言之就是判断它是否会移回到原来的位置。

移动顺序由一个字符串表示。每一个动作都是由一个字符来表示的。机器人有效的动作有 R（右），L（左），U（上）和 D（下）。输出应为 true 或 false，表示机器人移动路线是否成圈。
https://leetcode-cn.com/problems/judge-route-circle/description/
 */

func judgeCircle(moves string) bool {
	buffer := make([]int, 'U'+1)
	for _,c := range moves {
		buffer[c]++
	}
	return buffer['U'] == buffer['D'] && buffer['R'] == buffer['L']
}

func TestJudgeCircle(){
	tests := []struct{
		moves string
		result bool
	}{
		{
			"LLLL",
			false,
		},
		{
			"RRRR",
			false,
		},
		{
			"UUUU",
			false,
		},
		{
			"DDDD",
			false,
		},
		{
			"LLRR",
			true,
		},
		{
			"UUDD",
			true,
		},
		{
			"LRDU",
			true,
		},
	}
	for _,test := range tests{
		fmt.Println(judgeCircle(test.moves) == test.result)
	}
}
