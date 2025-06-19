package dsa

type Stack[T comparable] []T

func (s *Stack[T]) size() int {
	return len(*s)
}

func (s *Stack[T]) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack[T]) push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) isEmpty() bool {
	return len(*s) != 0
}
