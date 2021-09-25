package brackets

type Stack struct {
	elems []int
}

func New() *Stack {
	var stack = Stack{nil}
	return &stack
}

func (stack *Stack) Push(elem int) {
	if stack.elems == nil {
		stack.elems = []int{}
	}
	stack.elems = append(stack.elems, elem)
}

func (stack *Stack) Pop() int {
	x := stack.elems[len(stack.elems)-1]
	stack.elems = stack.elems[0 : len(stack.elems)-1]
	return x
}

func (stack *Stack) Size() int {
	return len(stack.elems)
}

func Bracket(str string) (bool, error) {
	_map := map[string]int{
		"(": 1,
		")": -1,
		"{": 2,
		"}": -2,
		"[": 3,
		"]": -3,
	}
	var stack = *New()
	for i := 0; i < len(str); i++ {
		j := _map[string(str[i])]
		if j > 0 {
			stack.Push(j)
		} else {
			if stack.Size() == 0 {
				return false, nil
			}
			result := j + stack.Pop()
			if result != 0 {
				return false, nil
			}
		}
	}
	return stack.Size() == 0, nil
}
