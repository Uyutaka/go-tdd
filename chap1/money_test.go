package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := New(5)
		assert.Equal(t, New(10), five.times(2))
		assert.Equal(t, New(15), five.times(3))
	})
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
