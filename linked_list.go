package main

import (
	"fmt"
	"strings"
)

type LinkedListInterface[T comparable] interface {
	Add(value *T)
	AddFirst(value *T)
	Remove(value *T)
	TraverseForward() <-chan T
	TraverseBackward() <-chan T
	String() string
}

var _ LinkedListInterface[int] = (*DoublyLinkedList[int])(nil)

type Node[T comparable] struct {
	Value      T
	Prev, Next *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Size int
}

func (list *DoublyLinkedList[T]) TraverseForward() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		current := list.Head
		for current != nil {
			ch <- current.Value
			current = current.Next
		}
	}()
	return ch
}

func (list *DoublyLinkedList[T]) TraverseBackward() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		current := list.Tail
		for current != nil {
			ch <- current.Value
			current = current.Prev
		}
	}()
	return ch
}

func (list *DoublyLinkedList[T]) Add(value *T) {
	newNode := &Node[T]{Value: *value}
	if list.Size == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}
	list.Size++
}

func (list *DoublyLinkedList[T]) AddFirst(value *T) {
	newNode := &Node[T]{Value: *value}
	if list.Size == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
	list.Size++
}

func (list *DoublyLinkedList[T]) Remove(value *T) {
	current := list.Head
	for current != nil {
		if current.Value == *value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				list.Head = current.Next

			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				list.Tail = current.Prev
			}
			list.Size--
			return
		}
		current = current.Next
	}
}

func (list *DoublyLinkedList[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	for e := list.Head; e != nil; e = e.Next {
		fmt.Fprintf(&sb, "%v", e.Value)
		if e.Next != nil {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
