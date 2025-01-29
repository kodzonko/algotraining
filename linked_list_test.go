package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	val1 := 1
	val2 := 2
	val3 := 3
	list.Add(&val1)
	list.Add(&val2)
	list.Add(&val3)

	if list.Size != 3 {
		t.Errorf("expected size 3, got %d", list.Size)
	}

	if list.Head.Value != 1 {
		t.Errorf("expected head value 1, got %d", list.Head.Value)
	}

	if list.Tail.Value != 3 {
		t.Errorf("expected tail value 3, got %d", list.Tail.Value)
	}
}

func TestAddToFront(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddFirst(3)

	if list.Size != 3 {
		t.Errorf("expected size 3, got %d", list.Size)
	}

	if list.Head.Value != 3 {
		t.Errorf("expected head value 3, got %d", list.Head.Value)
	}

	if list.Tail.Value != 1 {
		t.Errorf("expected tail value 1, got %d", list.Tail.Value)
	}
}

func TestRemove(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	val1 := 1
	val2 := 2
	val3 := 3

	// Setup list
	list.Add(&val1)
	list.Add(&val2)
	list.Add(&val3)

	// Remove middle element
	list.Remove(&val2)
	if list.Size != 2 {
		t.Errorf("expected size 2, got %d", list.Size)
	}
	if list.Head.Next != list.Tail {
		t.Error("links not updated after middle removal")
	}

	// Remove head
	list.Remove(&val1)
	if list.Size != 1 {
		t.Errorf("expected size 1, got %d", list.Size)
	}
	if list.Head.Value != 3 {
		t.Errorf("expected head value 3, got %d", list.Head.Value)
	}

	// Remove last element
	list.Remove(&val3)
	if list.Size != 0 {
		t.Errorf("expected size 0, got %d", list.Size)
	}
	if list.Head != nil || list.Tail != nil {
		t.Error("head and tail should be nil after removing all elements")
	}
}

func TestString(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	val1 := 1
	val2 := 2
	val3 := 3
	list.Add(&val1)
	list.Add(&val2)
	list.Add(&val3)

	expected := "[1, 2, 3]"
	if list.String() != expected {
		t.Errorf("expected %s, got %s", expected, list.String())
	}
}
