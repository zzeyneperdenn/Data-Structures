package internal

import "fmt"

type Node struct {
	value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(value int) *LinkedList {
	return &LinkedList{
		Head: &Node{
			value: value,
		},
	}
}

func (list *LinkedList) Push(value int) {
	// create new node
	new_node := &Node{
		value: value,
	}
	// append head to new nodes next

	new_node.Next = list.Head
	list.Head = new_node

	// head should be new node
}

func (list *LinkedList) InsertAfter(n *Node, value int) {
	newNode := &Node{
		value: value,
	}
	temp := n.Next
	n.Next = newNode
	newNode.Next = temp
}

func (list *LinkedList) Append(value int) {
	newNode := &Node{
		value: value,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {

		node := list.Head
		for node.Next != nil {
			node = node.Next
		}

		node.Next = newNode
	}
}

func (list *LinkedList) Print() {

	// 5 -> 6 -> 10

	for list.Head != nil {
		fmt.Printf("%d ->", list.Head.value)
		list.Head = list.Head.Next
	}

}
