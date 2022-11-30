package list

// Stack implements a generic LIFO data structure.
type Stack[E comparable] struct {
	items    *node[E]
	Size     uint64
	Capacity uint64
}

type node[E comparable] struct {
	value E
	next  *node[E]
}

// NewStack returns a pointer to a new Stack object.
func NewStack[E comparable](cap uint64) *Stack[E] {
	return &Stack[E]{
		items:    nil,
		Size:     0,
		Capacity: cap,
	}
}

// Push adds a new item to the stack. No-op and returns false if stack is full.
func (s *Stack[E]) Push(item E) (ok bool) {
	if s.Size == s.Capacity {
		return false
	}

	new := &node[E]{value: item}
	if s.items != nil {
		new.next = s.items
	}

	s.items = new
	s.Size++
	return true
}

// Pop removes an item from the stack. No-op and returns false if the stack is empty.
func (s *Stack[E]) Pop() (item E, ok bool) {
	if s.Size == 0 {
		return item, false
	}

	newHead := s.items.next
	item = s.items.value
	s.items.next = nil
	s.items = newHead

	s.Size--
	return item, true
}

// Peek returns but does not remove the item from the stack.
// No-op and returns false if the stack is empty.
func (s *Stack[E]) Peek() (item E, ok bool) {
	if s.Size == 0 {
		return item, false
	}

	item = s.items.value
	return item, true
}
