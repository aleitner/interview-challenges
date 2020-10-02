package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t*testing.T) {
	queue := &Queue{}
	queue.Insert("hello")
	queue.Insert("world")

	fmt.Printf("%+v", queue.current)

	data, err := queue.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "hello", data)

	data, err = queue.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "world", data)

	data, err = queue.Pop()
	assert.Error(t, err)
}
