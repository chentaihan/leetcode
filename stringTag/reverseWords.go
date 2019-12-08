package stringTag

import "fmt"

/**
151. 翻转字符串里的单词
示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 

说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseWords1(s string) string {
	buf := make([]byte, len(s))
	count := 0
	end := -1
	for index := len(s) - 1; index >= 0; index-- {
		if end == -1 && s[index] == ' ' {
			continue
		}
		if end == -1 {
			end = index
		} else if s[index] == ' ' {
			if count == 0 {
				copy(buf, s[index+1:end+1])
				count = end - index
			} else {
				buf[count] = ' '
				copy(buf[count+1:], s[index+1:end+1])
				count += end - index + 1
			}
			end = -1
		}
	}
	if end > -1 {
		if count == 0 {
			copy(buf, s[:end+1])
			count = end +1
		} else {
			buf[count] = ' '
			copy(buf[count+1:], s[:end+1])
			count += end + 2
		}
	}
	return string(buf[:count])
}

func TestreverseWords1() {
	tests := []struct{
		source string
		dest string
	}{
		{
			"abc",
			"abc",
		},
		{
			"   abc ",
			"abc",
		},
		{
			"abc edf qwe",
			"qwe edf abc",
		},
		{
			"abc edf",
			"edf abc",
		},
		{
			"      abc edf      ",
			"edf abc",
		},
		{
			"    abc edf qwe    ",
			"qwe edf abc",
		},
		{
			"abc      edf    qwe    123! ",
			"123! qwe edf abc",
		},
		{
			"    abc edf qwe    123!        ",
			"123! qwe edf abc",
		},
	}

	for _,test := range tests{
		dest := reverseWords1(test.source)
		if dest != test.dest {
			fmt.Println(dest,"|", test.dest)
		}else{
			fmt.Println(true)
		}
	}
}
