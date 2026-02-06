// queue.go 
// Maxie Machado 
// Linked List 
// Febuary 05, 2026 

package ds 

import "errors"

type Queue struct {
    list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	q.list.Insert(value) 
}

// remove the head
func (q *Queue) Pop() (string, error) {
	if q.list.Size == 0 {
		return "", errors.New("Pop: Queue is Empty") 
	}
	old := q.list.Head 
	val := old.data 

	q.list.Head = old.Next 
	old.Next = nil 
	q.list.Size-- 

	if q.list.Size == 0 {
		q.list.Tail = nil 
	}
	return val, nil 
}
