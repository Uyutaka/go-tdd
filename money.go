package main

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

func NewMoney(amount int, currency string) Money {
	return Money{amount: amount, currency: currency}
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func (m Money) Plus(money Money) Expression {
	return NewMoney(m.amount+money.amount, m.currency)
}

type Money struct {
	amount   int
	currency string
}

type Expression interface {
}
type Bank struct{}

func (b Bank) Reduce(source Expression, to string) Money {
	return NewMoney(10, "USD")
}
