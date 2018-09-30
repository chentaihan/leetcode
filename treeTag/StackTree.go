package treeTag


type Stack struct {
	buffer []*TreeNode
	size   int
}

func NewStackTree() Stack {
	return Stack{}
}

func (this *Stack) Push(x *TreeNode) {
	if len(this.buffer) == this.size {
		this.buffer = append(this.buffer, x)
	} else {
		this.buffer[this.size] = x
	}
	this.size++
}

func (this *Stack) Pop() *TreeNode {
	if this.size == 0 {
		return nil
	}
	val := this.buffer[this.size-1]
	this.size--
	return val
}

func (this *Stack) Top() *TreeNode {
	if this.size == 0 {
		return nil
	}
	return this.buffer[this.size-1]
}

func (this *Stack) Empty() bool {
	return this.size == 0
}

