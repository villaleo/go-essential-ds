package list

type Queue[E any] struct {
	items    []E
	Capacity uint64
}

func New[E any](cap uint64) *Queue[E] {
	return &Queue[E]{
		items:    make([]E, 0, cap),
		Capacity: cap,
	}
}

func (q *Queue[E]) Enqueue(item E) error {
	if len(q.items) == int(q.Capacity) {
		return FullQueueException{}
	}

	q.items = append(q.items, item)
	return nil
}

func (q *Queue[E]) Dequeue() (*E, error) {
	if len(q.items) == 0 {
		return nil, EmptyQueueException{}
	}

	val := q.items[0]
	q.items = q.items[1:]
	return &val, nil
}

func (q *Queue[E]) Peek() (*E, error) {
	if len(q.items) == 0 {
		return nil, EmptyQueueException{}
	}

	return &q.items[0], nil
}

type EmptyQueueException struct{}

func (e EmptyQueueException) Error() string {
	return "Queue is empty"
}

type FullQueueException struct{}

func (e FullQueueException) Error() string {
	return "Queue is full"
}
