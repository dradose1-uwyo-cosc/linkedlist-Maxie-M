// main.go 
// Maxie Machado 
// Linked List 
// Febuary 05, 2026 

package main 

import (
	"fmt" 
	"hw-linked/ds"
)

func main() {
	fmt.Println("=== LinkedList Demo ===")
	var l ds.LinkedList
	l.Insert("a")
	l.Insert("b")
	l.Insert("c")
	l.PrintList()

	fmt.Println("\nInsertAt(1, \"x\")")
	_ = l.InsertAt(1, "x") 
	l.PrintList()

	fmt.Println("\nRemove(\"b\")")
	_ = l.Remove("b") 
	l.PrintList()

	fmt.Println("\nReverse()")
	l.Reverse()
	l.PrintList()

	fmt.Println("\n=== Stack Demo ===") 
	var s ds.Stack 
	s.Push("one") 
	s.Push("two")
	s.Push("three") 

	for {
		v, ok := s.Pop()
		if !ok {
			break 
		}
		fmt.Println("popped:", v)
	}

	fmt.Println("\n=== Queue Demo ===")
	var q ds.Queue 
	q.Push("first") 
	q.Push("second") 
	q.Push("third") 
	for i := 0; i < 4; i++ {
		v,err := q.Pop()
		if err != nil {
			fmt.Println("pop error:", err)
			break
		}
		fmt.Println("popped:", v)
	}
}
