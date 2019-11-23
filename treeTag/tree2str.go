package treeTag

import (
	"bytes"
	"fmt"
	"strconv"
)

/**
606. 根据二叉树创建字符串
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

输入: 二叉树: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。
示例 2:

输入: 二叉树: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-string-from-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func tree2str(t *TreeNode) string {
	var buf bytes.Buffer
	_tree2str(t, &buf)
	return buf.String()
}

func _tree2str(t *TreeNode, buf *bytes.Buffer) {
	if t != nil {
		buf.WriteString(strconv.Itoa(t.Val))
		if t.Left != nil || t.Right != nil {
			buf.WriteString("(")
			_tree2str(t.Left, buf)
			buf.WriteString(")")
			if t.Right != nil {
				buf.WriteString("(")
				_tree2str(t.Right, buf)
				buf.WriteString(")")
			}
		}
	}
}

func Testtree2str() {
	tests := []struct {
		num    []int
		result string
	}{
		{
			[]int{1, 2, 3, 4},
			"1(2(4))(3)",
		},
		{
			[]int{1, 2, 3,-1, 4},
			"1(2()(4))(3)",
		},
	}

	for _, test := range tests {
		root := ArrayToTree(test.num)
		result := tree2str(root)
		fmt.Println(result == test.result)
	}
}
