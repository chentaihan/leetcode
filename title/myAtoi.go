package title

import "fmt"

func myAtoi(str string) int {
	result := 0
	symbol := 1
	isFirstChar := true
	for _, c := range str {
		if isFirstChar {
			if c == ' ' {
				continue
			}
			if c == '-' {
				symbol = -1
				isFirstChar = false
				continue
			}
			if c == '+' {
				isFirstChar = false
				continue
			}
		}

		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
			if result > 1<<31-1{
				break
			}
			isFirstChar = false
		} else {
			break
		}
	}
	result *= symbol
	if result > 1<<31-1 {
		return 1<<31 - 1
	}
	if result < -1<<31 {
		return -1 << 31
	}
	return result
}

func TestMyAtoi() {
	fmt.Println(myAtoi(" ") == 0)
	fmt.Println(myAtoi("  +12345") == 12345)
	fmt.Println(myAtoi("  -12345") == -12345)
	fmt.Println(myAtoi("  12345") == 12345)
	fmt.Println(myAtoi("12345") == 12345)
	fmt.Println(myAtoi("-12345") == -12345)
	fmt.Println(myAtoi("--12345") == 0)
	fmt.Println(myAtoi("+12345") == 12345)
	fmt.Println(myAtoi("12345qa") == 12345)
	fmt.Println(myAtoi("  12345wsxdr erg-12+12345") == 12345)
	fmt.Println(myAtoi("12345 with words") == 12345)
	fmt.Println(myAtoi("words and 987") == 0)
	fmt.Println(myAtoi("-words and 987") == 0)
	fmt.Println(myAtoi("--words and 987") == 0)
	fmt.Println(myAtoi("-91283472332") == -2147483648)
	fmt.Println(myAtoi("2147483648") == 2147483647)
	fmt.Println(myAtoi("9223372036854775808") == 2147483647)
}
