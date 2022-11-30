package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	queueCapacity uint64 = 100
)

func TestNewQueue(t *testing.T) {
	got := NewQueue[int](queueCapacity)

	assert.NotNil(t, got)
	assert.Equal(t, queueCapacity, got.capacity)
	assert.Equal(t, uint64(0), got.Size)
}

func TestEnqueue_HappyPath(t *testing.T) {
	queue := NewQueue[int](queueCapacity)
	assert.NotNil(t, queue)

	numIterations := 50
	for i := 1; i <= numIterations; i++ {
		// Enqueue 50 items
		ok := queue.Enqueue(i)
		assert.True(t, ok)
	}

	assert.Equal(t, uint64(numIterations), queue.Size)
}

func TestEnqueue_FullQueue(t *testing.T) {
	queue := NewQueue[int](queueCapacity)
	assert.NotNil(t, queue)

	for i := 1; i <= int(queueCapacity); i++ {
		// Enqueue 100 items (capacity limit)
		ok := queue.Enqueue(i)
		assert.True(t, ok)
	}

	// Queue is full. Attempt to add another item
	ok := queue.Enqueue(101)
	assert.False(t, ok)
	assert.Equal(t, queueCapacity, queue.Size)
}

func TestDequeue_HappyPath(t *testing.T) {
	queue := NewQueue[int](queueCapacity)
	assert.NotNil(t, queue)

	// Add 50 items to the queue
	numIterations := 50
	for i := 1; i <= numIterations; i++ {
		ok := queue.Enqueue(i)
		assert.True(t, ok)
	}

	// Dequeue each item
	for i := 1; i <= numIterations; i++ {
		val, ok := queue.Dequeue()
		assert.True(t, ok)
		assert.Equal(t, i, val)
	}

	assert.Equal(t, uint64(0), queue.Size)
}

func TestDequeue_EmptyQueue(t *testing.T) {
	queue := NewQueue[int](queueCapacity)
	assert.NotNil(t, queue)

	// Queue is empty. Attempt to remove another item
	_, ok := queue.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, uint64(0), queue.Size)
}

func TestPeek_HappyPath(t *testing.T) {
	queue := NewQueue[int](queueCapacity)
	assert.NotNil(t, queue)

	// Add 100 items to the queue
	numIterations := 100
	for i := 1; i <= numIterations; i++ {
		ok := queue.Enqueue(i)
		assert.True(t, ok)
	}

	// Peek the first item in queue (should be 1)
	val, ok := queue.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
}

func TestPeek_EmptyQueue(t *testing.T) {
	queue := NewQueue[int](queueCapacity)
	assert.NotNil(t, queue)

	// Attempt to peek at empty queue
	_, ok := queue.Peek()
	assert.False(t, ok)
	assert.Equal(t, uint64(0), queue.Size)
}
