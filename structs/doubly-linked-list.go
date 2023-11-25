package structs

import "fmt"

type DoublyListNode struct {
	data int
	prev *DoublyListNode
	next *DoublyListNode
}
type DoublyList struct {
	head *DoublyListNode
	tail *DoublyListNode
}

func (doublyList *DoublyList) PrintDoublyList() {
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

func (doublyList *DoublyList) PushBack(data int) {
	last := &DoublyListNode{data: data, prev: doublyList.tail, next: nil}

	if doublyList.head == nil {
		doublyList.head = last
		doublyList.tail = last
		return
	}

	doublyList.tail.next = last
	doublyList.tail = last
}

func (doublyList *DoublyList) PushFront(data int) {
	oldHead := doublyList.head
	head := &DoublyListNode{data: data, prev: nil, next: oldHead}
	doublyList.head = head
	oldHead.prev = head
}

func (doublyList *DoublyList) PopFront() (value int) {
	value = doublyList.head.data
	doublyList.head = doublyList.head.next
	return value
}

func (doublyList *DoublyList) PopBack() (value int) {
	value = doublyList.tail.data
	doublyList.tail = doublyList.tail.prev
	doublyList.tail.next = nil
	return value
}

func (doublyList *DoublyList) Insert(data, index int) {
	if index == 0 {
		doublyList.PushFront(data)
		return
	}
	find := doublyList.head
	for i := 0; find.next != nil && i < index-1; i++ {
		find = find.next
	}
	newNode := &DoublyListNode{data: data, prev: find, next: find.next}
	find.next = newNode
	newNode.next.prev = newNode
}
