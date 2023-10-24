package main

type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money{amount: amount, currency: "USD"}}
}

type Franc struct {
	Money
}

func NewFranc(amount int) Franc {
	return Franc{Money{amount: amount, currency: "CHF"}}
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

func NewMoney(amount int, currency string) Money {
	return Money{amount: amount, currency: currency}
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.Currency() == other.Currency()
}

func DollarToMoney(d Dollar) Money {
	return NewMoney(d.amount, "USD")
}

func FrancToMoney(f Franc) Money {
	return NewMoney(f.amount, "CHF")
}

func (m Money) Currency() string {
	return m.currency
}

type Money struct {
	amount   int
	currency string
}
