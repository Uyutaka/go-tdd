package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), five.times(2))
		assert.Equal(t, NewDollar(15), five.times(3))
	})
}

func TestEquals(t *testing.T) {
	t.Run("should return true if two Dollars have same amount", func(t *testing.T) {
		dollar := NewDollar(5)
		assert.True(t, dollar.equals(NewDollar(5)))
	})
	t.Run("should return false if two Dollars have different amount", func(t *testing.T) {
		dollar := NewDollar(5)
		assert.False(t, dollar.equals(NewDollar(7)))
	})
	t.Run("should return true if two Francs have same amount", func(t *testing.T) {
		franc := NewFranc(5)
		assert.True(t, franc.equals(NewFranc(5)))
	})
	t.Run("should return false if two Francs have different amount", func(t *testing.T) {
		franc := NewFranc(5)
		assert.False(t, franc.equals(NewFranc(7)))
	})

	t.Run("should return false when Franc and Dollar are compared", func(t *testing.T) {
		franc := NewFranc(5)
		assert.False(t, franc.equals(NewDollar(5)))
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := NewFranc(5)
		assert.Equal(t, NewFranc(10), five.times(2))
		assert.Equal(t, NewFranc(15), five.times(3))
	})
}
