package main

// NodeKey base type for cache values
type NodeKey interface{}

// Node base type for list elements
type Node struct {
	Key    NodeKey
	MapKey string
	Prev   *Node
	Next   *Node
}

// List is a circular doubly linked list
type List struct {
	Head   *Node
	Length int
}

//NewList initializes an empty list
func NewList() *List {
	return &List{nil, 0}
}

// Insert adds a new key to the end
func (l *List) Insert(key NodeKey, mapKey string) *Node {
	node := &Node{Key: key, MapKey: mapKey}
	l.Length++

	if l.Head == nil {
		l.Head = node
		node.Prev = node
		node.Next = node
		return node
	}

	p := l.Head.Prev
	p.Next = node
	node.Prev = p
	node.Next = l.Head
	l.Head.Prev = node

	return node
}

// RemoveNode removes a specific Node from the list
func (l *List) RemoveNode(node *Node) *Node {
	if l.Head == nil {
		return nil
	}

	l.Length--

	if l.Head == node {
		oldHead := l.Head
		l.Head = l.Head.Next
		l.Head.Prev = oldHead.Prev
		return node
	}

	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev

	return node
}

// RemoveTail convenience method for remove the tail
func (l *List) RemoveTail() *Node {
	if l.Head == nil {
		return nil
	}

	return l.RemoveNode(l.Head.Prev)
}
