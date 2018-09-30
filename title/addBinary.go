package title

import "fmt"

func addBinary(a string, b string) string {
	aLen := len(a) - 1
	bLen := len(b) - 1
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	buffer := make([]byte, maxLen+1)
	prevSum := uint8(0)
	for aLen >= 0 && bLen >= 0 {
		curSum := (a[aLen] - '0' + b[bLen] - '0') + prevSum
		switch curSum {
		case 0:
			buffer[maxLen] = '0'
			prevSum = 0
		case 1:
			buffer[maxLen] = '1'
			prevSum = 0
		case 2:
			buffer[maxLen] = '0'
			prevSum = 1
		case 3:
			buffer[maxLen] = '1'
			prevSum = 1
		}
		maxLen--
		aLen--
		bLen--
	}
	for aLen >= 0 {
		curSum := a[aLen] - '0' + prevSum
		switch curSum {
		case 0:
			buffer[maxLen] = '0'
			prevSum = 0
		case 1:
			buffer[maxLen] = '1'
			prevSum = 0
		case 2:
			buffer[maxLen] = '0'
			prevSum = 1
		case 3:
			buffer[maxLen] = '1'
			prevSum = 1
		}
		maxLen--
		aLen--
	}
	for bLen >= 0 {
		curSum := b[bLen] - '0' + prevSum
		switch curSum {
		case 0:
			buffer[maxLen] = '0'
			prevSum = 0
		case 1:
			buffer[maxLen] = '1'
			prevSum = 0
		case 2:
			buffer[maxLen] = '0'
			prevSum = 1
		case 3:
			buffer[maxLen] = '1'
			prevSum = 1
		}
		maxLen--
		bLen--
	}
	if prevSum == 1 {
		buffer[maxLen] = '1'
		return string(buffer)
	}
	return string(buffer[1:])
}

func TestAddBinary() {
	fmt.Println(addBinary("0", "1") == "1")
	fmt.Println(addBinary("1", "0") == "1")
	fmt.Println(addBinary("0", "0") == "0")
	fmt.Println(addBinary("1", "1") == "10")
	fmt.Println(addBinary("1", "11") == "100")
	fmt.Println(addBinary("11", "11") == "110")
	fmt.Println(addBinary("11", "111") == "1010")
	fmt.Println(addBinary("11", "1111") == "10010")
	fmt.Println(addBinary("101", "1010") == "1111")
	fmt.Println(addBinary("101", "1011") == "10000")
	fmt.Println(addBinary("101", "11011") == "100000")
	fmt.Println(addBinary("101", "111011") == "1000000")
	fmt.Println(addBinary("101", "10111011") == "11000000")
}
