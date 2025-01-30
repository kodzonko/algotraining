package main

import (
	"testing"
)

type listFixture struct {
	val1, val2, val3 int
	list             *DoublyLinkedList[int]
}

func setupList() *listFixture {
	f := &listFixture{
		val1: 1,
		val2: 2,
		val3: 3,
		list: &DoublyLinkedList[int]{},
	}
	return f
}

func TestAppend(t *testing.T) {
	f := setupList()
	f.list.Add(&f.val1)
	f.list.Add(&f.val2)
	f.list.Add(&f.val3)

	if f.list.Size != 3 {
		t.Errorf("expected size 3, got %d", f.list.Size)
	}

	if f.list.Head.Value != 1 {
		t.Errorf("expected head value 1, got %d", f.list.Head.Value)
	}

	if f.list.Tail.Value != 3 {
		t.Errorf("expected tail value 3, got %d", f.list.Tail.Value)
	}
}

func TestAddToFront(t *testing.T) {
	f := setupList()

	f.list.AddFirst(&f.val3)
	f.list.AddFirst(&f.val2)
	f.list.AddFirst(&f.val1)

	if f.list.Size != 3 {
		t.Errorf("expected size 3, got %d", f.list.Size)
	}

	if f.list.Head.Value != f.val1 {
		t.Errorf("expected head value %d, got %d", f.val1, f.list.Head.Value)
	}

	if f.list.Tail.Value != f.val3 {
		t.Errorf("expected tail value %d, got %d", f.val3, f.list.Tail.Value)
	}

	if f.list.Head.Next.Value != f.val2 {
		t.Errorf("expected second node value %d, got %d", f.val2, f.list.Head.Next.Value)
	}
}

func TestRemove(t *testing.T) {
	f := setupList()
	f.list.Add(&f.val1)
	f.list.Add(&f.val2)
	f.list.Add(&f.val3)

	f.list.Remove(&f.val2)
	if f.list.Size != 2 {
		t.Errorf("expected size 2, got %d", f.list.Size)
	}
	if f.list.Head.Next != f.list.Tail {
		t.Error("links not updated after middle removal")
	}

	f.list.Remove(&f.val1)
	if f.list.Size != 1 {
		t.Errorf("expected size 1, got %d", f.list.Size)
	}
	if f.list.Head.Value != 3 {
		t.Errorf("expected head value 3, got %d", f.list.Head.Value)
	}

	f.list.Remove(&f.val3)
	if f.list.Size != 0 {
		t.Errorf("expected size 0, got %d", f.list.Size)
	}
	if f.list.Head != nil || f.list.Tail != nil {
		t.Error("head and tail should be nil after removing all elements")
	}
}

func TestString(t *testing.T) {
	f := setupList()
	f.list.Add(&f.val1)
	f.list.Add(&f.val2)
	f.list.Add(&f.val3)

	expected := "[1, 2, 3]"
	if f.list.String() != expected {
		t.Errorf("expected %s, got %s", expected, f.list.String())
	}
}

func readChannel[T any](ch <-chan T) []T {
	var result []T
	for v := range ch {
		result = append(result, v)
	}
	return result
}

func TestTraverseForward(t *testing.T) {
	f := setupList()

	result := readChannel(f.list.TraverseForward())
	if len(result) != 0 {
		t.Errorf("expected empty list, got %v", result)
	}

	f.list.Add(&f.val1)
	f.list.Add(&f.val2)
	f.list.Add(&f.val3)

	expected := []int{1, 2, 3}
	result = readChannel(f.list.TraverseForward())

	if len(result) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

func TestTraverseBackward(t *testing.T) {
	f := setupList()

	f.list.Add(&f.val1)
	f.list.Add(&f.val2)
	f.list.Add(&f.val3)

	expected := []int{3, 2, 1}
	result := readChannel(f.list.TraverseBackward())

	if len(result) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}
