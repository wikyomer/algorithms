package structs

type Stack[TData any] struct {
	head List[TData]
}

func (stack *Stack[TData]) Push(data TData) {
	stack.head.PushFront(data)
}

func (stack *Stack[TData]) Pop() (value TData) {
	return stack.head.PopFront()
}

func (stack *Stack[TData]) Print() {
	stack.head.PrintList()
}

func (stack *Stack[TData]) IsEmpty() bool {
	return stack.head.IsEmpty()
}

func (stack *Stack[TData]) Pick() TData {
	return stack.head.Pick()
}
