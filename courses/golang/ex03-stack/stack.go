package stack

type Stack struct {
	elems []int
}

func New() *Stack {
	var stack Stack = Stack{nil}
	return &stack
}

func (stack *Stack) Push(elems int) {
	if stack.elems == nil {
		stack.elems = []int{}
	}
	stack.elems = append(stack.elems, elems)
}

func (stack *Stack) Pop() int {
	c := stack.elems[len(stack.elems)-1]
	stack.elems = stack.elems[0 : len(stack.elems)-1]
	return c
}

func (stack *Stack) Size() int {
	return len(stack.elems)
}
