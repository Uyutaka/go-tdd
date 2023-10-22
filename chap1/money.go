package main

type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money{amount}}
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
	return Franc{Money{amount}}
}

func (m Money) Times(multiplier int) Money {
	return Money{m.amount * multiplier}
}

func (f *Franc) equals(anotherMoney any) bool {
	franc, ok := anotherMoney.(Franc)
	if !ok {
		return false
	}
	return f.amount == franc.amount
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.GetType() == other.GetType()
}

func (d Dollar) GetType() string {
	return "Dollar"
}
func (f Franc) GetType() string {
	return "Franc"
}
func (m Money) GetType() string {
	return "Money"
}

type Money struct {
	amount int
}
