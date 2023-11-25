package structs

type Queue struct {
	head DoublyList
}

func (queue *Queue) Push(data int) {
	queue.head.PushBack(data)
}

func (queue *Queue) Pop() (value int) {
	return queue.head.PopFront()
}

func (queue *Queue) Print() {
	queue.head.PrintDoublyList()
}
