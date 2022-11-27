package list

// Queue implements a generic FIFO data structure.
type Queue[E any] struct {
	items    []E
	capacity uint64
	Size     uint64
}

// New returns a pointer to a new Queue object.
func New[E any](cap uint64) *Queue[E] {
	return &Queue[E]{
		items:    make([]E, 0, cap),
		capacity: cap,
		Size:     0,
	}
}

// Enqueue adds an item to the queue. Returns false if queue is full.
func (q *Queue[E]) Enqueue(item E) (ok bool) {
	if len(q.items) == int(q.capacity) {
		return false
	}

	q.items = append(q.items, item)
	q.Size++
	return true
}

// Dequeue removes an item from the queue. Returns the item and false if queue is empty.
func (q *Queue[E]) Dequeue() (out E, ok bool) {
	if len(q.items) == 0 {
		return out, false
	}

	val := q.items[0]
	q.items = q.items[1:]
	q.Size--
	return val, true
}

// Peek returns but does not remove the item from the queue. Returns the item and false if queue is empty.
func (q *Queue[E]) Peek() (out E, ok bool) {
	if len(q.items) == 0 {
		return out, false
	}

	return q.items[0], true
}
