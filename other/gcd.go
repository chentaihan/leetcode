package other

import "fmt"

/**
最大公约数
 */
func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}

func TestGCD() {
	tests := []struct {
		x int
		y int
		z int
	}{
		{
			21,
			18,
			3,
		},
		{
			8,
			4,
			4,
		},
		{
			8,
			0,
			8,
		},
		{
			0,
			4,
			4,
		},
	}

	for _, test := range tests {
		z := GCD(test.x, test.y)
		fmt.Println(z == test.z)
	}
}
