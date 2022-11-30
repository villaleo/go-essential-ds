package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	stackCapacity uint64 = 100
	zero          uint64 = 0
)

func TestNewStack(t *testing.T) {
	var capacity uint64 = 100
	got := NewStack[int](capacity)

	assert.NotNil(t, got)
	assert.Equal(t, capacity, got.Capacity)
	assert.Equal(t, zero, got.Size)
}

func TestPush(t *testing.T) {
	var (
		expectedSize  uint64 = 50
		numIterations int    = 50
	)

	got := NewStack[int](stackCapacity)
	assert.NotNil(t, got)

	var ok bool
	for i := 1; i <= numIterations; i++ {
		// Assert all pushes are successful
		ok = got.Push(i)
		assert.True(t, ok)
	}

	assert.Equal(t, expectedSize, got.Size)
}

func TestPush_FullStack(t *testing.T) {
	var numIterations int = 100
	got := NewStack[int](stackCapacity)
	assert.NotNil(t, got)

	var ok bool
	for i := 1; i <= numIterations; i++ {
		ok = got.Push(i)
		assert.True(t, ok)
	}

	// Stack is full. Try to add one more item
	ok = got.Push(101)
	assert.False(t, ok)
	assert.Equal(t, uint64(numIterations), got.Size)
}

func TestPop(t *testing.T) {
	var numIterations = 50
	got := NewStack[int](stackCapacity)
	assert.NotNil(t, got)

	// Insert 50 items
	var ok bool
	for i := 1; i <= numIterations; i++ {
		ok = got.Push(i)
		assert.True(t, ok)
	}

	// Remove all items
	for i := 1; i <= numIterations; i++ {
		val, ok := got.Pop()
		assert.Equal(t, val, numIterations-i+1)
		assert.True(t, ok)
	}

	assert.Equal(t, zero, got.Size)
}

func TestPop_EmptyStack(t *testing.T) {
	got := NewStack[int](stackCapacity)
	assert.NotNil(t, got)

	// Attempt to pop from empty stack
	_, ok := got.Pop()
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	got := NewStack[int](stackCapacity)
	assert.NotNil(t, got)

	// Insert some items
	numItems := 35
	for i := 1; i <= numItems; i++ {
		ok := got.Push(i)
		assert.True(t, ok)
	}

	// Peek at the last item in the stack
	val, ok := got.Peek()
	assert.True(t, ok)
	assert.Equal(t, numItems, val)
}

func TestPeek_EmptyStack(t *testing.T) {
	got := NewStack[int](stackCapacity)
	assert.NotNil(t, got)

	_, ok := got.Peek()
	assert.False(t, ok)
}
