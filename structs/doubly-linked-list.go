package structs

import "fmt"

type DoublyListNode[TData any] struct {
	data TData
	prev *DoublyListNode[TData]
	next *DoublyListNode[TData]
}
type DoublyList[TData any] struct {
	head *DoublyListNode[TData]
	tail *DoublyListNode[TData]
}

func (doublyList *DoublyList[TData]) PrintDoublyList() {
	if doublyList.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := doublyList.head
	for current != nil {
		fmt.Printf("%v\n", current.data)
		current = current.next
	}
}

func (doublyList *DoublyList[TData]) PushBack(data TData) {
	last := &DoublyListNode[TData]{data: data, prev: doublyList.tail, next: nil}

	if doublyList.head == nil {
		doublyList.head = last
		doublyList.tail = last
		return
	}

	doublyList.tail.next = last
	doublyList.tail = last
}

func (doublyList *DoublyList[TData]) PushFront(data TData) {
	oldHead := doublyList.head
	head := &DoublyListNode[TData]{data: data, prev: nil, next: oldHead}
	doublyList.head = head
	oldHead.prev = head
}

func (doublyList *DoublyList[TData]) PopFront() (value TData) {
	value = doublyList.head.data
	doublyList.head = doublyList.head.next
	return value
}

func (doublyList *DoublyList[TData]) PopBack() (value TData) {
	value = doublyList.tail.data
	doublyList.tail = doublyList.tail.prev
	doublyList.tail.next = nil
	return value
}

func (doublyList *DoublyList[TData]) Insert(data TData, index int) {
	if index == 0 {
		doublyList.PushFront(data)
		return
	}
	find := doublyList.head
	for i := 0; find.next != nil && i < index-1; i++ {
		find = find.next
	}
	newNode := &DoublyListNode[TData]{data: data, prev: find, next: find.next}
	find.next = newNode
	newNode.next.prev = newNode
}

func (doublyList *DoublyList[TData]) IsEmpty() bool {
	return doublyList.head == nil
}
