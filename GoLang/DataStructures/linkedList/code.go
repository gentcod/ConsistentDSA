package linkedList

import (
	"fmt"
)

func main() {
	fmt.Print("Hello")
}

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T any] struct {
	count int
	head *Node[T]
	tail *Node[T]
}

func New[T any](values ...T) *SinglyLinkedList[T] {
	list := SinglyLinkedList[T]{}
	list.Add(values...)
	return &list
}

func (s *SinglyLinkedList[T]) IsEmpty() bool {
	if s.head == nil {
		return true
	}
	return false
}

func (s *SinglyLinkedList[T]) Size() int {
	return s.count
}

func (s *SinglyLinkedList[T]) Add(values ...T) {
	for _,value := range values {
		s.InsertAt(value, s.Size())
	}
}

func (s *SinglyLinkedList[T]) InsertAt(value T, index int) error {
	if index < 0 || index > s.count {
		return nil
	}

	element := &Node[T]{Data: value}

	if s.IsEmpty() {
		s.head = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}

	currentIndex := 0
	current := s.head

	for currentIndex != index {
		current = current.Next
		currentIndex += 1
	}

	prevNode := current
	nextNode := current.Next

	prevNode.Next = element
	element.Next = nextNode

	return nil
}