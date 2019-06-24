package main

import (
	"testing"
)

func TestListInsert(t *testing.T) {
	list := NewList()

	for i := 0; i < 10; i++ {
		list.Insert(i, "")
	}

	head := list.Head
	for i := 0; i < 10; i++ {
		if head.Key != i {
			t.Errorf("Expected %d == %d", head.Key, i)
		}

		head = head.Next
	}

	if head != list.Head {
		t.Errorf("Expected iterator to equal Head")
	}
}

func TestInsertSingle(t *testing.T) {
	list := NewList()
	list.Insert(0, "")

	head := list.Head

	if list.Head != head.Next || list.Head != head.Prev {
		t.Fail()
	}

}

func TestListRemoveNodeEmpty(t *testing.T) {
	list := NewList()

	removed := list.RemoveNode(&Node{5, "", nil, nil})
	if removed != nil {
		t.Fail()
	}
}

func TestListRemoveHead(t *testing.T) {
	list := NewList()

	toRemove := list.Insert(1, "")
	list.Insert(2, "")

	list.RemoveNode(toRemove)
	if list.Head.Key != 2 {
		t.Errorf("Expected head to contain 2")
	}
}

func TestListRemoveMiddle(t *testing.T) {
	list := NewList()

	list.Insert(1, "")
	toRemove := list.Insert(2, "")
	list.Insert(3, "")

	list.RemoveNode(toRemove)

	if list.Head.Key != 1 {
		t.Errorf("Expected head to contain 1")
	}

	if list.Head.Next.Key != 3 {
		t.Errorf("Expected value after head to contain 3")
	}
}

func TestListRemoveEnd(t *testing.T) {
	list := NewList()

	list.Insert(1, "")
	list.Insert(2, "")
	toRemove := list.Insert(3, "")

	list.RemoveNode(toRemove)

	if list.Head.Key != 1 {
		t.Errorf("Expected head to contain 1")
	}

	if list.Head.Prev.Key != 2 {
		t.Errorf("Expected value before head to contain 3")
	}
}

func TestListRemoveLeastFrequent(t *testing.T) {
	list := NewList()

	list.Insert(1, "")
	list.Insert(2, "")
	list.Insert(3, "")

	removed := list.RemoveTail()
	if removed.Key != 3 {
		t.Fail()
	}

	if list.Length != 2 {
		t.Fail()
	}
}
