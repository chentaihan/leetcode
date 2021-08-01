package stackTag

import "fmt"

/**
当你依次访问完一串页面 a-b-c 之后，点击浏览器的后退按钮，就可以查看之前浏览过的页面 b 和 a。当你后退到页面 a，点击前进按钮，就可以重新查看页面 b 和 c。但是，如果你后退到页面 b 后，点击了新的页面 d，那就无法再通过前进、后退功能查看页面 c 了。
*/

type chromeStack struct {
	stack []int
	index int
}

func NewChromeStack() *chromeStack{
	return &chromeStack{
		stack: []int{},
		index: -1,
	}
}

func (cs *chromeStack) Push(val int) {
	if cs.index == len(cs.stack)-1 {
		cs.stack = append(cs.stack, val)
	} else {
		cs.stack[cs.index+1] = val
	}
	cs.index++
}

func (cs *chromeStack) Front() int {
	if cs.index == len(cs.stack)-1 {
		return -1
	}
	cs.index++
	return cs.stack[cs.index]
}

func (cs *chromeStack) Back() int {
	if cs.index == 0 {
		return -1
	}
	cs.index--
	return cs.stack[cs.index]
}

func TestchromeStack() {
	cs := NewChromeStack()
	cs.Push(1)
	cs.Push(2)
	cs.Push(3)
	cs.Back()
	cs.Push(4)
	cs.Push(5)
	cs.Back()
	fmt.Println(cs.stack)
}
