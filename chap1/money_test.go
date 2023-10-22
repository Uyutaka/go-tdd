package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	five := New(5)
	product := five.times(2)
	assert.Equal(t, 10, product.amount)
	product = five.times(3)
	assert.Equal(t, 15, product.amount)
}

func TestEquality(t *testing.T) {
	t.Run("should return true if two Dollars have same amount", func(t *testing.T) {
		dollar := New(5)
		assert.True(t, dollar.equals(New(5)))
	})
	t.Run("should return false if two Dollars have different amount", func(t *testing.T) {
		dollar := New(5)
		assert.False(t, dollar.equals(New(7)))
	})
}
