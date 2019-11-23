package mapTag

/**
1002. 查找常用字符
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。

示例 1：

输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：["cool","lock","cook"]
输出：["c","o"]
 

提示：

1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-common-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func commonChars(A []string) []string {
	constStr := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	arrA := make([][26]int, len(A))
	for index, item := range A {
		for _, c := range item {
			arrA[index][c-'a']++
		}
	}
	var result  []string
	for j := 0; j < 26; j++ {
		min := 1<<63 - 1
		for i := 0; i < len(arrA); i++ {
			if arrA[i][j] <= 0 {
				min = 0
				break
			}
			if arrA[i][j] < min {
				min = arrA[i][j]
			}
		}
		for i := 0; i < min; i++ {
			result = append(result, constStr[j])
		}
	}
	return result
}
