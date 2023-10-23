package main

type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money{amount: amount, currency: "USD"}}
}

func (d *Dollar) equals(anotherMoney any) bool {
	dollar, ok := anotherMoney.(Dollar)
	if !ok {
		return false
	}
	return d.amount == dollar.amount
}

type Franc struct {
	Money
}

func NewFranc(amount int) Franc {
	return Franc{Money{amount: amount, currency: "CHF"}}
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier}
}

func (f *Franc) equals(anotherMoney any) bool {
	franc, ok := anotherMoney.(Franc)
	if !ok {
		return false
	}
	return f.amount == franc.amount
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.Currency() == other.Currency()
}

func (m Money) Currency() string {
	return m.currency
}

type Money struct {
	amount   int
	currency string
}
