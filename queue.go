package dsago

type Queue[T comparable] []T

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Dequeue() T {
	frontElement := (*q)[1]
	*q = (*q)[1:]
	return frontElement
}

func (q *Queue[T]) Enqueue(element T) {
	*q = append(*q, element)
}

func (q *Queue[T]) Peek() T {
	return (*q)[1]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) != 0
}
