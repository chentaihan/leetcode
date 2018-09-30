package title

import "fmt"

func removeElement(nums []int, val int) int {
	start, end, len := 0, len(nums)-1, len(nums)
	for start <= end {
		for nums[start] != val {
			start++
			if start > end {
				break
			}
		}
		for nums[end] == val {
			end--
			len--
			if start > end {
				break
			}
		}
		if start > end {
			break
		}
		nums[start] = nums[end]
		start++
		end--
		len--
	}
	return len
}

func TestRemoveElement() {
	fmt.Println(removeElement([]int{}, 0) == 0)
	fmt.Println(removeElement([]int{1}, 1) == 0)
	fmt.Println(removeElement([]int{1}, 0) == 1)
	fmt.Println(removeElement([]int{1, 2}, 1) == 1)
	fmt.Println(removeElement([]int{1, 2}, 2) == 1)
	fmt.Println(removeElement([]int{1, 2}, 3) == 2)
	fmt.Println(removeElement([]int{1, 2, 3}, 3) == 2)
	fmt.Println(removeElement([]int{1, 2, 3}, 2) == 2)
	fmt.Println(removeElement([]int{1, 2, 3}, 1) == 2)
	fmt.Println(removeElement([]int{1, 2, 3}, 4) == 3)
	fmt.Println(removeElement([]int{1, 1, 2, 3}, 1) == 2)
	fmt.Println(removeElement([]int{1, 1, 2, 1, 3}, 1) == 2)
	fmt.Println(removeElement([]int{1, 1, 2, 1, 3, 1}, 1) == 2)
	fmt.Println(removeElement([]int{1, 1, 2, 1, 3, 1, 1}, 1) == 2)
	fmt.Println(removeElement([]int{1, 1, 2, 3, 1, 1}, 1) == 2)
	fmt.Println(removeElement([]int{1, 1, 2, 2, 3, 1, 1}, 1) == 3)
	fmt.Println(removeElement([]int{1, 1, 1, 2, 3, 1, 1}, 1) == 2)
	fmt.Println(removeElement([]int{1, 1, 1, 2, 3, 1, 1,3}, 1) == 3)
	fmt.Println(removeElement([]int{1, 2, 1, 3, 4, 5, 1,6}, 1) == 5)
}
