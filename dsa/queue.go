package dsa

type Queue[T comparable] []T

func (q *Queue[T]) size() int {
	return len(*q)
}

func (q *Queue[T]) dequeue() T {
	frontElement := (*q)[1]
	*q = (*q)[1:]
	return frontElement
}

func (q *Queue[T]) enqueue(element T) {
	*q = append(*q, element)
}

func (q *Queue[T]) peek() T {
	return (*q)[1]
}

func (q *Queue[T]) isEmpty() bool {
	return len(*q) != 0
}
