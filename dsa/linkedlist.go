package dsa

// Singly Linked List
type node[T comparable] struct {
	element T
	next    *node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
}

func (s *SinglyLinkedList[T]) len() int {
	counter := 0
	currentNode := s.head
	for currentNode != nil {
		counter++
		currentNode = currentNode.next
	}
	return counter
}

func (s *SinglyLinkedList[T]) index(inputNode *node[T]) int {
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

func swap[T comparable](firstNode, secondNode *node[T]) {
	firstNode.element, secondNode.element = secondNode.element, firstNode.element
}

func (s *SinglyLinkedList[T]) previousNode(inputNode *node[T]) *node[T] {
	currentNode := s.head
	for range s.index(inputNode) - 1 {
		currentNode = currentNode.next
	}
	return currentNode
}

// Insert
func (s *SinglyLinkedList[T]) insertAtHead(element T) {
	newNode := &node[T]{element: element, next: s.head}
	s.head = newNode
	if s.tail == nil {
		s.tail = newNode
	}
}

func (s *SinglyLinkedList[T]) insertAtTail(element T) {
	newNode := &node[T]{element: element}
	if s.tail != nil {
		s.tail.next = newNode
	}
	s.tail = newNode
	if s.head == nil {
		s.head = newNode
	}
}

func (s *SinglyLinkedList[T]) insertAtIndex(element T, index int) {
	if index < 0 {
		panic("index cannot be negative")
	}
	if index == 0 {
		s.insertAtHead(element)
		return
	}
	if index > s.len() {
		panic("index out of bounds")
	}
	newNode := &node[T]{
		element: element,
		next:    nil,
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
func (s *SinglyLinkedList[T]) removeHead() {
	if s.len() == 0 {
		panic("Empty linked list")
	}
	s.head = s.head.next
}

func (s *SinglyLinkedList[T]) removeTail() {
	currentNode := s.head
	nextNode := s.head.next
	for range s.len() - 2 {
		currentNode = nextNode
	}
	s.tail = currentNode
}

func (s *SinglyLinkedList[T]) removeAtIndex(index int) {
	if index < 0 {
		panic("index cannot be negative")
	}
	if index == 0 {
		s.removeHead()
		return
	}
	if index > s.len() {
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

func (s *SinglyLinkedList[T]) get(index int) *node[T] {
	currentNode := s.head
	for range index {
		currentNode = currentNode.next
	}
	return s.head
}

func (s *SinglyLinkedList[T]) find(element T) int {
	currentNode := s.head
	index := 0
	for {
		if currentNode.element == element {
			return index
		}
		currentNode = currentNode.next
		index++
	}
}

func (s *SinglyLinkedList[T]) reverse() {
	currentNode := s.head
	var previousNode *node[T]
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode

	}
}
