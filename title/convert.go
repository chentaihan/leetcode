package title

import "fmt"

func convert1(s string, numRows int) string {
	fmt.Println(s)
	if numRows <= 1 {
		return s
	}
	zLen := (numRows - 1 ) * 2
	zCount := len(s) / zLen
	index := 0
	buffer := make([]byte, len(s))
	for i := 0; i < zLen; i++ {
		for j := 0; j < zCount; j++ {
			buffer[index] = s[i+j*zLen]
			fmt.Print(i+j*zLen, " ")
			index++
		}
		if i+zCount*zLen < len(s) {
			buffer[index] = s[i+zCount*zLen]
			fmt.Print(i+zCount*zLen, " ")
			index++
		}
	}
	fmt.Println()
	fmt.Println(string(buffer))
	return string(buffer)
}

func convert(s string, numRows int) string {
	fmt.Println(s)
	if numRows <= 1 {
		return s
	}
	vSize := (numRows - 1 ) * 2
	vCount := (len(s)-1)/vSize + 1
	bufferIndex, lineIndex := 0, 0
	buffer := make([]byte, len(s))
	for vIndex := 0; vIndex < vSize; vIndex++ {
		for i := 0; i < vCount; i++ {
			vFirst := i * vSize

			if vFirst+lineIndex < len(s) {
				buffer[bufferIndex] = s[vFirst+lineIndex]
				bufferIndex++
			}

			if lineIndex > 0 && lineIndex < numRows-1 {
				if vFirst+vSize-lineIndex < len(s) {
					buffer[bufferIndex] = s[vFirst+vSize-lineIndex]
					bufferIndex++
				}
			}
			lineIndex++
		}
	}

	fmt.Println()
	fmt.Println(string(buffer))
	return string(buffer)
}

func TestConvert() {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}
