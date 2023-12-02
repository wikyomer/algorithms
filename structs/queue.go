package structs

type Queue[TData any] struct {
	list DoublyList[TData]
}

func (queue *Queue[TData]) Push(data TData) {
	queue.list.PushBack(data)
}

func (queue *Queue[TData]) Pop() (value TData) {
	return queue.list.PopFront()
}

func (queue *Queue[TData]) Print() {
	queue.list.PrintDoublyList()
}

func (queue *Queue[TData]) IsEmpty() bool {
	return queue.list.IsEmpty()
}
