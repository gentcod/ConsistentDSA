package linkedList

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T any] struct {
	count int
	head *Node[T]
	tail *Node[T]
}

// New helps to initialize the LinkedList with specific values
func New[T any](values ...T) *SinglyLinkedList[T] {
	list := SinglyLinkedList[T]{}
	
	list.Add(values...)
	return &list
}

// NewEmpty helps to initialize an empty LinkedList
func NewEmpty() *SinglyLinkedList[any] {
	list := SinglyLinkedList[any]{}
	
	return &list
}


// IsEmpty checks if the LinkedList is empty
func (s *SinglyLinkedList[T]) IsEmpty() bool {
	return s.head == nil
}

// Size returns the length or size of the LinkedList
func (s *SinglyLinkedList[T]) Size() int {
	return s.count
}

// Add adds a new node to the head, replacing the previous head of the LinkedList.
func (s *SinglyLinkedList[T]) Add(values ...T) {
	for _,value := range values {
		// fmt.Println(value)
		node := &Node[T]{Data: value}
		node.Next = s.head
		s.head = node
		s.count++

		if s.count == 1 {
			s.tail = node.Next
		}
	}
}

// InsertAt takes a value and an index and adds the node to the list based on the index.
func (s *SinglyLinkedList[T]) InsertAt(value T, index int) error {
	if index < 0 || index > s.count {
		return errors.New("index is out of bound")
	}

	if index == 0 {
		s.Add(value)
		return nil
	}

	element := &Node[T]{Data: value}

	currentIndex := index
	current := s.head

	for currentIndex > 1 {
		current = current.Next
		currentIndex--
	}

	prev := current
	nextNode := current.Next

	prev.Next = element
	element.Next = nextNode
	s.count++

	return nil
}

// RemoveAt removes a node in a particular index
func (s *SinglyLinkedList[T]) RemoveAt(index int) error {

	if index < 0 || index > s.count || index == s.count {
		return errors.New("index is out of bound")
	}

	if index == 0 {
		nextNode := s.head.Next
		s.head = nil
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
	data.Next = nil
	data = nil
	s.count--

	return nil
}

// Diplay iterates through the list and shows the data in the LinkedList
func (s *SinglyLinkedList[T]) Diplay() {
	data := []any {}

	current := s.head
	prev := &Node[T]{}

	for i := s.count; i > 0; i-- {
		if current == s.head {
			data = append(data, fmt.Sprintf("[Head: %v]", current.Data))
		} else if current == s.tail {
			data = append(data, fmt.Sprintf("[Tail: %v]", current.Data))
		} else {
			data = append(data, fmt.Sprintf("[Data: %v]", current.Data))
		}

		prev = current
		current = prev.Next
	}

	fmt.Println(data)
}