package title

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	index := 0
	array := [19]int{}
	for x != 0 {
		array[index] = x % 10
		index++
		x /= 10
	}
	start, end := 0, index -1
	for start < end{
		if array[start] != array[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func TestIsPalindrome(){
	fmt.Println(isPalindrome(-1) == false)
	fmt.Println(isPalindrome(0) == true)
	fmt.Println(isPalindrome(1) == true)
	fmt.Println(isPalindrome(12) == false)
	fmt.Println(isPalindrome(121) == true)
	fmt.Println(isPalindrome(1221) == true)
	fmt.Println(isPalindrome(123321) == true)
	fmt.Println(isPalindrome(1234321) == true)
	fmt.Println(isPalindrome(123454321) == true)
	fmt.Println(isPalindrome(12345654321) == true)
	fmt.Println(isPalindrome(1234567890987654321) == true)
	fmt.Println(isPalindrome(123234321) == false)
}