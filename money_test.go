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
		bank := NewBank()
		reduced := bank.Reduce(sum, "USD")
		assert.Equal(t, NewMoney(10, "USD"), reduced)
	})
	t.Run("plust should return sum", func(t *testing.T) {
		five := NewMoney(5, "USD")
		result := five.Plus(five)
		sum, ok := result.(Sum)
		if !ok {
			t.Errorf("result is not Sum")
		}
		assert.Equal(t, five, sum.augend)
		assert.Equal(t, five, sum.addend)
	})
}
func TestReduce(t *testing.T) {
	t.Run("return correct total dollars", func(t *testing.T) {
		sum := NewSum(NewMoney(3, "USD"), NewMoney(4, "USD"))
		bank := NewBank()
		result := bank.Reduce(sum, "USD")
		assert.Equal(t, NewMoney(7, "USD"), result)
	})
	t.Run("reduce money", func(t *testing.T) {
		bank := NewBank()
		result := bank.Reduce(NewMoney(3, "USD"), "USD")
		assert.Equal(t, NewMoney(3, "USD"), result)
	})
}
func TestAddRate(t *testing.T) {
	t.Run("reduce money different currency", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(NewMoney(2, "CHF"), "USD")
		assert.Equal(t, NewMoney(1, "USD"), result)
	})
}

func TestRate(t *testing.T) {
	t.Run("rate should return 1 when two currencies are same", func(t *testing.T) {
		bank := NewBank()
		assert.Equal(t, 1, bank.Rate("USD", "USD"))
	})
}

func TestMixedAddition(t *testing.T) {
	t.Run("should return correct total dollars", func(t *testing.T) {
		fiveBucks := NewMoney(5, "USD")
		tenFrancs := NewMoney(10, "CHF")
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
		assert.Equal(t, NewMoney(10, "USD"), result)
	})
	t.Run("should return correct total dollars with Money", func(t *testing.T) {
		fiveBucks := NewMoney(5, "USD")
		tenFrancs := NewMoney(10, "CHF")
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		sum := NewSum(fiveBucks, tenFrancs).Plus(fiveBucks)
		result := bank.Reduce(sum, "USD")
		assert.Equal(t, NewMoney(15, "USD"), result)
	})
}
func TestSumTimes(t *testing.T) {
	t.Run("should return correct total dollars using sum.Times()", func(t *testing.T) {
		fiveBucks := NewMoney(5, "USD")
		tenFrancs := NewMoney(10, "CHF")
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		sum := NewSum(fiveBucks, tenFrancs).Times(2)
		result := bank.Reduce(sum, "USD")
		assert.Equal(t, NewMoney(20, "USD"), result)
	})
}
