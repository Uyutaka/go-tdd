package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), NewDollar(five.Times(2).amount))
		assert.Equal(t, NewDollar(15), NewDollar(five.Times(3).amount))
	})
	t.Run("should return same amount if Dollar creates from factory method", func(t *testing.T) {
		five := NewDollar(5)
		assert.Equal(t, NewDollar(10), NewDollar(five.Times(2).amount))
		assert.Equal(t, NewDollar(15), NewDollar(five.Times(3).amount))
	})
}

func TestEquals(t *testing.T) {
	t.Run("should return true if two Dollars have same amount", func(t *testing.T) {
		dollar := DollarToMoney(NewDollar(5))
		assert.True(t, NewDollar(5).Equals(dollar))
	})
	t.Run("should return false if two Dollars have different amount", func(t *testing.T) {
		dollar := DollarToMoney(NewDollar(7))
		assert.False(t, NewDollar(5).Equals(dollar))
	})
	t.Run("should return true if two Francs have same amount", func(t *testing.T) {
		franc := FrancToMoney(NewFranc(5))
		assert.True(t, NewFranc(5).Equals(franc))
	})
	t.Run("should return false if two Francs have different amount", func(t *testing.T) {
		franc := FrancToMoney(NewFranc(7))
		assert.False(t, NewFranc(5).Equals(franc))
	})

	t.Run("should return false when Franc and Dollar are compared", func(t *testing.T) {
		dollar := DollarToMoney(NewDollar(5))
		assert.False(t, NewFranc(7).Equals(dollar))
	})

	t.Run("should return true when currency", func(t *testing.T) {
		money := NewMoney(5, "USD")
		dollar := DollarToMoney(NewDollar(5))
		assert.True(t, money.Equals(dollar))
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("should return same amount if Dollar multiplies by times()", func(t *testing.T) {
		five := NewFranc(5)
		assert.Equal(t, NewFranc(10), NewFranc(five.Times(2).amount))
		assert.Equal(t, NewFranc(15), NewFranc(five.Times(3).amount))
	})
}

func TestCurrency(t *testing.T) {
	t.Run("should return correct currency name", func(t *testing.T) {
		assert.Equal(t, "CHF", NewFranc(4).Currency())
		assert.Equal(t, "USD", NewDollar(4).Currency())
	})
}
