package dsa

// Singly Linked List
type singlyNode[T comparable] struct {
	value T
	next  *singlyNode[T]
}

type SinglyLinkedList[T comparable] struct {
	head *singlyNode[T]
	tail *singlyNode[T]
}

func (s *SinglyLinkedList[T]) Len() int {
	counter := 0
	currentNode := s.head
	for currentNode != nil {
		counter++
		currentNode = currentNode.next
	}
	return counter
}

func (s *SinglyLinkedList[T]) Index(inputNode *singlyNode[T]) int {
	counter := 0
	currentNode := s.head
	for currentNode != nil {
		if currentNode == inputNode {
			return counter
		}
		counter++
		currentNode = currentNode.next
	}
	return counter
}

func Swap[T comparable](firstNode, secondNode *singlyNode[T]) {
	firstNode.value, secondNode.value = secondNode.value, firstNode.value
}

func (s *SinglyLinkedList[T]) PreviousNode(inputNode *singlyNode[T]) *singlyNode[T] {
	currentNode := s.head
	for range s.Index(inputNode) - 1 {
		currentNode = currentNode.next
	}
	return currentNode
}

// Insert
func (s *SinglyLinkedList[T]) InsertAtHead(value T) {
	newNode := &singlyNode[T]{value: value, next: s.head}
	s.head = newNode
	if s.tail == nil {
		s.tail = newNode
	}
}

func (s *SinglyLinkedList[T]) InsertAtTail(value T) {
	newNode := &singlyNode[T]{value: value}
	if s.tail != nil {
		s.tail.next = newNode
	}
	s.tail = newNode
	if s.head == nil {
		s.head = newNode
	}
}

func (s *SinglyLinkedList[T]) InsertAtIndex(value T, index int) {
	if index < 0 {
		panic("index cannot be negative")
	}
	if index == 0 {
		s.InsertAtHead(value)
		return
	}
	if index > s.Len() {
		panic("index out of bounds")
	}
	newNode := &singlyNode[T]{
		value: value,
		next:  nil,
	}
	currentNode := s.head
	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			panic("index out of bounds")
		}
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next = newNode
}

// Remove
func (s *SinglyLinkedList[T]) RemoveHead() {
	if s.Len() == 0 {
		panic("Empty linked list")
	}
	s.head = s.head.next
}

func (s *SinglyLinkedList[T]) RemoveTail() {
	currentNode := s.head
	nextNode := s.head.next
	for range s.Len() - 2 {
		currentNode = nextNode
	}
	s.tail = currentNode
}

func (s *SinglyLinkedList[T]) RemoveAtIndex(index int) {
	if index < 0 {
		panic("index cannot be negative")
	}
	if index == 0 {
		s.RemoveHead()
		return
	}
	if index > s.Len() {
		panic("index out of bounds")
	}
	currentNode := s.head
	nextNode := s.head.next
	for range index - 1 {
		currentNode = nextNode
		if currentNode == nil {
			panic("index out of bounds")
		}
	}
	currentNode.next = nextNode.next.next
}

func (s *SinglyLinkedList[T]) Get(index int) *singlyNode[T] {
	currentNode := s.head
	for range index {
		currentNode = currentNode.next
	}
	return s.head
}

func (s *SinglyLinkedList[T]) Find(value T) int {
	currentNode := s.head
	index := 0
	for {
		if currentNode.value == value {
			return index
		}
		currentNode = currentNode.next
		index++
	}
}

func (s *SinglyLinkedList[T]) Reverse() {
	currentNode := s.head
	var previousNode *singlyNode[T]
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
}
