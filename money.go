package main

func (m Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

func NewMoney(amount int, currency string) Money {
	return Money{amount: amount, currency: currency}
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func (m Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}

type Money struct {
	amount   int
	currency string
}

type Expression interface {
	Plus(addend Expression) Expression
	Reduce(bank Bank, to string) Money
}
type Bank struct {
	rates map[Pair]int
}

func NewBank() Bank {
	return Bank{rates: make(map[Pair]int)}
}

func (b Bank) Rate(from string, to string) int {
	if from == to {
		return 1
	}
	return b.rates[NewPair(from, to)]
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}

func (b Bank) AddRate(from string, to string, rate int) {
	pair := NewPair(from, to)
	b.rates[pair] = rate
}

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend Expression, addend Expression) Sum {
	return Sum{augend: augend, addend: addend}
}

func (m Sum) Reduce(bank Bank, to string) Money {
	augend := m.augend.Reduce(bank, to)
	addend := m.addend.Reduce(bank, to)
	total := augend.amount + addend.amount
	return NewMoney(total, to)
}

type Pair struct {
	from string
	to   string
}

func (m Sum) Plus(addend Expression) Expression {
	// TODO temporary implementation
	return Sum{}
}

func NewPair(from string, to string) Pair {
	return Pair{from: from, to: to}
}

func (p Pair) equals(other Pair) bool {
	return p.from == other.from && p.to == other.to
}

func (p Pair) hashCode() int {
	return 0
}
