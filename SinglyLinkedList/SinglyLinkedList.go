package main

import (
	"errors"
	"fmt"
)

// Node represents a single node in the list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// SinglyLinkedList represents a singly linked list
type SinglyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Len  int
}

// NewSinglyLinkedList creates a new empty singly linked list
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// PushBack adds a new element to the end of the list
func (l *SinglyLinkedList[T]) PushBack(v T) {
	n := &Node[T]{Value: v}
	if l.Tail == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Len++
}

// PushFront adds a new element to the beginning of the list
func (l *SinglyLinkedList[T]) PushFront(v T) {
	n := &Node[T]{Value: v}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
	l.Len++
}

// PopFront removes and returns the first element of the list
func (l *SinglyLinkedList[T]) PopFront() (T, error) {
	var zero T
	if l.Head == nil {
		return zero, errors.New("empty list")
	}
	v := l.Head.Value
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	l.Len--
	return v, nil
}

// Print traverses and prints all elements in the list
func (l *SinglyLinkedList[T]) Print() {
	cur := l.Head
	for cur != nil {
		fmt.Print(cur.Value, " -> ")
		cur = cur.Next
	}
	fmt.Println("nil")
}

// Example usage
func main() {
	list := NewSinglyLinkedList[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushFront(0)
	list.Print() // Output: 0 -> 1 -> 2 -> nil

	val, _ := list.PopFront()
	fmt.Println("PopFront:", val)
	list.Print() // Output: 1 -> 2 -> nil
}
