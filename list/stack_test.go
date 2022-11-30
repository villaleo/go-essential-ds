package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	stackCapacity uint64 = 100
)

func TestNewStack(t *testing.T) {
	var capacity uint64 = 100
	stack := NewStack[int](capacity)

	assert.NotNil(t, stack)
	assert.Equal(t, capacity, stack.Capacity)
	assert.Equal(t, uint64(0), stack.Size)
}

func TestPush(t *testing.T) {
	stack := NewStack[int](stackCapacity)
	assert.NotNil(t, stack)

	numIterations := 50
	for i := 1; i <= numIterations; i++ {
		// Assert all pushes are successful
		ok := stack.Push(i)
		assert.True(t, ok)
	}

	assert.Equal(t, uint64(numIterations), stack.Size)
}

func TestPush_FullStack(t *testing.T) {
	stack := NewStack[int](stackCapacity)
	assert.NotNil(t, stack)

	numIterations := 100
	for i := 1; i <= numIterations; i++ {
		ok := stack.Push(i)
		assert.True(t, ok)
	}

	// Stack is full. Try to add one more item
	ok := stack.Push(101)
	assert.False(t, ok)
	assert.Equal(t, uint64(numIterations), stack.Size)
}

func TestPop(t *testing.T) {
	stack := NewStack[int](stackCapacity)
	assert.NotNil(t, stack)

	// Insert 50 items
	numIterations := 50
	for i := 1; i <= numIterations; i++ {
		ok := stack.Push(i)
		assert.True(t, ok)
	}

	// Remove all items
	for i := 1; i <= numIterations; i++ {
		val, ok := stack.Pop()
		assert.Equal(t, val, numIterations-i+1)
		assert.True(t, ok)
	}

	assert.Equal(t, uint64(0), stack.Size)
}

func TestPop_EmptyStack(t *testing.T) {
	stack := NewStack[int](stackCapacity)
	assert.NotNil(t, stack)

	// Attempt to pop from empty stack
	_, ok := stack.Pop()
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	stack := NewStack[int](stackCapacity)
	assert.NotNil(t, stack)

	// Insert some items
	numIterations := 35
	for i := 1; i <= numIterations; i++ {
		ok := stack.Push(i)
		assert.True(t, ok)
	}

	// Peek at the last item in the stack
	val, ok := stack.Peek()
	assert.True(t, ok)
	assert.Equal(t, numIterations, val)
}

func TestPeek_EmptyStack(t *testing.T) {
	stack := NewStack[int](stackCapacity)
	assert.NotNil(t, stack)

	// Attempt to peek at empty stack
	_, ok := stack.Peek()
	assert.False(t, ok)
	assert.Equal(t, uint64(0), stack.Size)
}
