package title

import (
	"fmt"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > (1<<31)-1 || result < -1<<31 {
		return 0
	}
	return result
}

func TestReverse() {
	fmt.Println((1<<31)-1)
	fmt.Println(reverse(123))
	fmt.Println(reverse(1230))
	fmt.Println(reverse(-1230))
}
