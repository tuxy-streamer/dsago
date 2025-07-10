package dsago

type Stack[T comparable] []T

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Pop() T {
	lastElement := (*s)[len(*s)-1]
	*s = (*s)[:s.Size()-1]
	return lastElement
}

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	return (*s)[s.Size()-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}
