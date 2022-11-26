package list

type Queue[E any] struct {
	items    []E
	capacity uint64
	Size     uint64
}

func New[E any](cap uint64) *Queue[E] {
	return &Queue[E]{
		items:    make([]E, 0, cap),
		capacity: cap,
		Size:     0,
	}
}

func (q *Queue[E]) Enqueue(item E) (ok bool) {
	if len(q.items) == int(q.capacity) {
		return false
	}

	q.items = append(q.items, item)
	q.Size++
	return true
}

func (q *Queue[E]) Dequeue() (out E, ok bool) {
	if len(q.items) == 0 {
		return out, false
	}

	val := q.items[0]
	q.items = q.items[1:]
	q.Size--
	return val, true
}

func (q *Queue[E]) Peek() (out E, ok bool) {
	if len(q.items) == 0 {
		return out, false
	}

	return q.items[0], true
}
