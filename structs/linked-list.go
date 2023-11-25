package structs

import "fmt"

type ListNode[TData any] struct {
	data TData
	next *ListNode[TData]
}

type List[TData any] struct {
	head *ListNode[TData]
}

func (list *List[TData]) PrintList() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("%v\n", current.data)
		current = current.next
	}
}

func (list *List[TData]) PushBack(data TData) {
	last := &ListNode[TData]{data: data, next: nil}

	if list.head == nil {
		list.head = last
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = last
}

func (list *List[TData]) Insert(data TData, index int) {
	if index == 0 {
		list.PushFront(data)
		return
	}
	find := list.head
	for i := 0; find.next != nil && i < index-1; i++ {
		find = find.next
	}
	newNode := &ListNode[TData]{data: data, next: find.next}
	find.next = newNode
}

func (list *List[TData]) PushFront(data TData) {
	head := &ListNode[TData]{data: data, next: list.head}
	list.head = head
}

func (list *List[TData]) PopFront() (value TData) {
	value = list.head.data
	list.head = list.head.next
	return value
}

func (list *List[TData]) IsEmpty() bool {
	return list.head == nil
}

func (list *List[TData]) Pick() TData {
	return list.head.data
}
