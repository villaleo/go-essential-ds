package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var cap uint64 = 10
	want := &Queue[int]{
		capacity: cap,
	}
	got := NewQueue[int](cap)

	assert.NotNil(t, got)
	assert.Equal(t, want.capacity, got.capacity)
	assert.Equal(t, got.items, make([]int, 0, cap))
}

func TestEnqueue_HappyPath(t *testing.T) {
	var cap uint64 = 3
	queue := NewQueue[int](cap)

	ok := queue.Enqueue(1)
	assert.True(t, ok)
	ok = queue.Enqueue(2)
	assert.True(t, ok)

	assert.Equal(t, queue.items, []int{1, 2})
}

func TestEnqueue_FullQueue(t *testing.T) {
	var cap uint64 = 1
	queue := NewQueue[int](cap)

	ok := queue.Enqueue(1)
	assert.True(t, ok)
	ok = queue.Enqueue(2)
	assert.False(t, ok)

	assert.Equal(t, queue.items, []int{1})
}

func TestDequeue_HappyPath(t *testing.T) {
	var cap uint64 = 3
	queue := NewQueue[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	queue.Enqueue(10)
	queue.Enqueue(20)

	val, ok := queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 10, val)
}

func TestDequeue_EmptyQueue(t *testing.T) {
	var cap uint64 = 3
	queue := NewQueue[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	_, ok := queue.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, queue.Size, uint64(0))
}

func TestPeek_HappyPath(t *testing.T) {
	var cap uint64 = 3
	queue := NewQueue[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	queue.Enqueue(10)
	queue.Enqueue(20)

	val, ok := queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, val, 10)
}

func TestPeek_EmptyQueue(t *testing.T) {
	var cap uint64 = 3
	queue := NewQueue[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	_, ok := queue.Dequeue()
	assert.False(t, ok)
}
