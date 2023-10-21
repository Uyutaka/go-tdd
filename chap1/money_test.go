package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testTimes(t *testing.T) {
	five := New(5)
	five.times(2)
	assert.Equal(t, 10, five.amount)
}
