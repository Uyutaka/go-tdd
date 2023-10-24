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

type Money struct {
	amount   int
	currency string
}
