// linkedlist.go 
// Maxie Machado 
// Linked List 
// Febuary 05, 2026 

package ds 

import (
	"errors"
	"fmt"
)

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}

// insert at the tail
func (l *LinkedList) Insert(value string) {
	n := newNode(value) 

	if l.Size == 0 {
		l.Head = n 
		l.Tail = n 
		l.Size = 1 
		return 
	}

	l.Tail.Next = n 
	l.Tail = n 
	l.Size++ 
}

//inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) InsertAt(position int, value string) error {
	if position < 0 || position > l.Size { 
		return fmt.Errorf("InsertAt: position %d out of range (0..%d)", position, l.Size)
	}
	n := newNode(value) 

	if l.Size == 0 {
		l.Head = n
		l.Tail = n 
		l.Size = 1 
		return nil 
	}

	if position == 0 {
		n.Next = l.Head 
		l.Head = n 
		l.Size++ 
		return nil 
	}

	if position == l.Size {
		l.Tail.Next = n 
		l.Tail = n 
		l.Size++ 
		return nil 
	}

	prev := l.Head 
	for i := 0; i < position-1; i++ {
		prev = prev.Next 
	}
	n.Next = prev.Next 
	prev.Next = n 
	l.Size++ 
	return nil 
}

// remove first occurrence of the value
func (l *LinkedList) Remove(value string) error {
	if l.Size == 0 {
		return errors.New("Remove: List is Empty") 
	}

	if l.Head.data == value {
		old := l.Head 
		l.Head = l.Head.Next 
		old.Next = nil 
		l.Size-- 

		if l.Size == 0 {
			l.Tail = nil 
		}
		return nil 
	}

	prev := l.Head 
	curr := l.Head.Next 
	
	for curr != nil {
		if curr.data == value {
			prev.Next = curr.Next 
			curr.Next = nil 
			l.Size-- 

			if prev.Next == nil {
				l.Tail = prev 
			}
			return nil 
		}
		prev = curr 
		curr = curr.Next 
	}
	return fmt.Errorf("Remove: value %q not found", value)
}

//remove all occurrences of a value
func (l *LinkedList) RemoveAll(value string) error {
	if l.Size == 0 {
		return errors.New("RemoveAll: List is Empty")
	}

	removedAny := false 

	for l.Head != nil && l.Head.data == value { 
		removedAny = true 
		old := l.Head 
		l.Head = l.Head.Next 
		old.Next = nil 
		l.Size-- 
	}

	if l.Head == nil {
		l.Tail = nil 

		if removedAny {
			return nil
		}
		return fmt.Errorf("RemoveAll: value %q not found", value)
	}

	prev := l.Head 
	curr := l.Head.Next 

	for curr != nil {
		if curr.data == value {
			removedAny = true 
			prev.Next = curr.Next 
			curr.Next = nil 
			l.Size--
			curr = prev.Next 
			continue 
		}
		prev = curr 
		curr = curr.Next 
	}

	t := l.Head 
	for t.Next != nil {
		t = t.Next 
	}
	l.Tail = t 

	if !removedAny {
		return fmt.Errorf("RemoveAll: value %q not found", value)
	}
	return nil 
}

// remove at a position, if index exists
func (l *LinkedList) RemoveAt(pos int) error {
	if l.Size == 0 {
		return errors.New("RemoveAt: List is Empty")
	}

	if pos < 0 || pos >= l.Size {
		return fmt.Errorf("RemoveAt: Position %d out of range (0..%d)", pos, l.Size-1)
	}

	if pos == 0 {
		old := l.Head 
		l.Head = l.Head.Next 
		old.Next = nil 
		l.Size-- 

		if l.Size == 0 {
			l.Tail = nil 
		}
		return nil 
	}
	prev := l.Head 

	for i := 0; i < pos-1; i++ {
		prev = prev.Next
	}

	target := prev.Next 
	prev.Next = target.Next 
	target.Next = nil 
	l.Size-- 

	if prev.Next == nil {
		l.Tail = prev 
	}
	return nil 
}

// checks if the linked list is empty
func (l *LinkedList) IsEmpty() bool {
	return l.Size == 0 
}

// get size of ll
func (l *LinkedList) GetSize() int {
	return l.Size 
}

//reverse the list
func (l *LinkedList) Reverse() {
	if l.Size <= 1 {
		return 
	}
	l.Tail = l.Head
	var prev *Node = nil 
	curr := l.Head 

	for curr != nil {
		next := curr.Next 
		curr.Next = prev 
		prev = curr 
		curr = next 
	}
	l.Head = prev 
}

//print the list 
func (l *LinkedList) PrintList() {
	if l.Size == 0 {
		fmt.Println("(empty)") 
		return 
	}
	curr := l.Head 

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.Next 
	}
}
