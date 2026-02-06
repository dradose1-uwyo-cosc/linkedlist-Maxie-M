// stack.go 
// Maxie Machado 
// Linked List 
// Febuary 05, 2026 

package ds 

type Stack struct {
    list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	n := newNode(value) 

	if s.list.Size == 0 {
		s.list.Head = n 
		s.list.Tail = n 
		s.list.Size = 1 
		return 
	}

	n.Next = s.list.Head 
	s.list.Head = n 
	s.list.Size++
}

// remove the head
func (s *Stack) Pop() (string, bool) {
	if s.list.Size == 0 {
		return "", false
	}

	old := s.list.Head 
	val := old.data 

	s.list.Head = old.Next 
	old.Next = nil 
	s.list.Size-- 

	if s.list.Size == 0 {
		s.list.Tail = nil 
	}
	return val, true 
}
