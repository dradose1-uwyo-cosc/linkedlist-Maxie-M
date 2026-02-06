// node.go 
// Maxie Machado 
// Linked List 
// Febuary 05, 2026 

package ds 

type Node struct { 
	data string 
	Next *Node 
}

func newNode(value string) *Node { 
	return &Node{data: value, Next: nil} 
}
