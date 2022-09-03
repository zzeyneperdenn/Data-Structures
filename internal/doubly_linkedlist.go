package internal

import "fmt"

type DNode struct {
	value    int
	Previous *DNode
	Next     *DNode
}

type DoublyLinkedList struct {
	Head *DNode
}

func NewDoublyLinkedList(value int) *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: &DNode{
			value: value,
		},
	}
}

func (list *DoublyLinkedList) Push(value int) {
	newNode := &DNode{
		value:    value,
		Previous: nil,
	}

	if list.Head != nil {
		newNode.Previous = newNode
	}

	newNode.Next = list.Head
	list.Head = newNode
}

func (list *DoublyLinkedList) DoublyAppend(value int) {
	newNode := &DNode{
		value: value,
		Next:  nil,
	}

	if list.Head.Next == nil {
		list.Head = newNode
	} else {
		node := list.Head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = newNode
	}
}

func (list *DoublyLinkedList) DoublyInsertAfter(prev *DNode, value int) {
	if prev == nil {
		fmt.Println("it should have a valid prev value")
		return
	}

	newNode := &DNode{
		value: value,
	}

	pn := prev.Next
	newNode.Next = pn
	prev.Next = newNode
	newNode.Previous = prev
	if newNode.Next != nil {
		newNode.Next.Previous = newNode

	}
}

func (list *DoublyLinkedList) Print() {
	for list.Head != nil {
		fmt.Printf("%d ->", list.Head.value)
		list.Head = list.Head.Next
	}
}
