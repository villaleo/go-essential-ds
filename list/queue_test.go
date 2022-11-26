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
	got := New[int](cap)

	assert.NotNil(t, got)
	assert.Equal(t, want.capacity, got.capacity)
	assert.Equal(t, got.items, make([]int, 0, cap))
}

func TestEnqueue_HappyPath(t *testing.T) {
	var cap uint64 = 3
	var err error
	queue := New[int](cap)

	err = queue.Enqueue(1)
	assert.Nil(t, err)
	err = queue.Enqueue(2)
	assert.Nil(t, err)

	assert.Equal(t, queue.items, []int{1, 2})
}

func TestEnqueue_FullQueue(t *testing.T) {
	var cap uint64 = 1
	var err error
	queue := New[int](cap)

	err = queue.Enqueue(1)
	assert.Nil(t, err)
	err = queue.Enqueue(2)
	assert.Equal(t, err, FullQueueException{})

	assert.Equal(t, queue.items, []int{1})
}

func TestDequeue_HappyPath(t *testing.T) {
	var cap uint64 = 3
	queue := New[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	queue.Enqueue(10)
	queue.Enqueue(20)

	val, err := queue.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 10, val)
}

func TestDequeue_EmptyQueue(t *testing.T) {
	var cap uint64 = 3
	queue := New[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	_, err := queue.Dequeue()
	assert.NotNil(t, err)
	assert.Equal(t, err, EmptyQueueException{})
	assert.Equal(t, queue.Size, uint64(0))
}

func TestPeek_HappyPath(t *testing.T) {
	var cap uint64 = 3
	queue := New[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	queue.Enqueue(10)
	queue.Enqueue(20)

	val, err := queue.Dequeue()
	assert.NotNil(t, val)
	assert.Equal(t, val, 10)
	assert.Nil(t, err)
}

func TestPeek_EmptyQueue(t *testing.T) {
	var cap uint64 = 3
	queue := New[int](cap)

	assert.NotNil(t, queue)
	assert.Equal(t, queue.capacity, cap)

	_, err := queue.Dequeue()
	assert.NotNil(t, err)
	assert.Equal(t, err, EmptyQueueException{})
}
