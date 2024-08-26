package linkedList

import (
	"errors"
	"fmt"
)

type DNode[T any] struct {
	Data T
	Prev *DNode[T]
	Next *DNode[T]
}

type DoublyLinkedList[T any] struct {
	count int
	head *DNode[T]
	tail *DNode[T]
}

// New helps to initialize the LinkedList with specific values
func NewDoubly[T any](values ...T) *DoublyLinkedList[T] {
	list := DoublyLinkedList[T]{}
	
	list.Add(values...)
	return &list
}

// NewEmpty helps to initialize an empty LinkedList
func NewEmptyDoubly() *DoublyLinkedList[any] {
	list := DoublyLinkedList[any]{}
	
	return &list
}


// IsEmpty checks if the LinkedList is empty
func (s *DoublyLinkedList[T]) IsEmpty() bool {
	return s.head == nil
}

// Size returns the length or size of the LinkedList
func (s *DoublyLinkedList[T]) Size() int {
	return s.count
}

// Add adds a new node to the head, replacing the previous head of the LinkedList.
func (s *DoublyLinkedList[T]) Add(values ...T) {
	for _,value := range values {
		node := &DNode[T]{Data: value}
		node.Next = s.head
		if s.count > 0 {
			node.Next.Prev = node
		}
		s.head = node
		
		if s.count == 1 {
			s.tail = node.Next
		}
		s.count++
	}
}

// InsertAt takes a value and an index and adds the node to the list based on the index.
func (s *DoublyLinkedList[T]) InsertAt(value T, index int) error {
	if index < 0 || index > s.count {
		return errors.New("index is out of bound")
	}

	if index == 0 {
		s.Add(value)
		return nil
	}

	element := &DNode[T]{Data: value}

	currentIndex := index
	current := s.head

	for currentIndex > 1 {
		current = current.Next
		currentIndex--
	}

	prev := current
	nextNode := current.Next

	prev.Next = element
	element.Prev = prev
	element.Next = nextNode
	nextNode.Prev = element
	s.count++

	return nil
}

// RemoveAt removes a node in a particular index
func (s *DoublyLinkedList[T]) RemoveAt(index int) error {

	if index < 0 || index > s.count || index == s.count {
		return errors.New("index is out of bound")
	}

	if index == 0 {
		nextNode := s.head.Next
		s.head = nil
		nextNode.Prev = nil
		s.head = nextNode
		s.count--
		return nil
	}

	currentIndex := index
	current := s.head

	for currentIndex > 1 {
		current = current.Next
		currentIndex--
	}

	data := current.Next
	new_next := data.Next

	current.Next = new_next
	new_next.Prev = current
	data.Next = nil
	data.Prev = nil
	data = nil
	s.count--

	return nil
}

// Diplay iterates through the list and shows the data in the LinkedList
func (s *DoublyLinkedList[T]) Diplay() {
	data := []any {}

	current := s.head
	prev := &DNode[T]{}

	for i := s.count; i > 0; i-- {
		if current == s.head {
			data = append(data, fmt.Sprintf("[Head: %v]", current.Data))
		} else if current == s.tail {
			data = append(data, fmt.Sprintf("[Tail: %v]", current.Data))
		} else {
			// data = append(data, fmt.Sprintf("[Data: %v]", current.Data))
			data = append(data, fmt.Sprintf("[Prev: %v]-[Data: %v]-[Next: %v]", current.Prev.Data, current.Data, current.Next.Data))
		}

		prev = current
		current = prev.Next
	}

	fmt.Println(data)
}