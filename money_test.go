package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := NewMoney(5, "USD")
		assert.Equal(t, NewMoney(10, "USD"), five.Times(2))
		assert.Equal(t, NewMoney(15, "USD"), five.Times(3))
	})
}

func TestEquals(t *testing.T) {
	t.Run("should return true if two Dollars have same amount", func(t *testing.T) {
		dollar := NewMoney(5, "USD")
		assert.True(t, NewMoney(5, "USD").Equals(dollar))
	})
	t.Run("should return false if two Dollars have different amount", func(t *testing.T) {
		dollar := NewMoney(7, "USD")
		assert.False(t, NewMoney(5, "USD").Equals(dollar))
	})

	t.Run("should return false when Franc and Dollar are compared", func(t *testing.T) {
		dollar := NewMoney(5, "USD")
		assert.False(t, NewMoney(7, "CHF").Equals(dollar))
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := NewMoney(5, "CHF")
		assert.Equal(t, NewMoney(10, "CHF"), five.Times(2))
		assert.Equal(t, NewMoney(15, "CHF"), five.Times(3))
	})
}

func TestCurrency(t *testing.T) {
	t.Run("should return correct currency name", func(t *testing.T) {
		assert.Equal(t, "CHF", NewMoney(10, "CHF").currency)
		assert.Equal(t, "USD", NewMoney(10, "USD").currency)
	})
}

func TestSimpleAddition(t *testing.T) {
	t.Run("should return $10 if $5 + $5", func(t *testing.T) {
		five := NewMoney(5, "USD")
		sum := five.Plus(five)
		bank := Bank{}
		reduced := bank.Reduce(sum, "USD")
		assert.Equal(t, NewMoney(10, "USD"), reduced)
	})
}
