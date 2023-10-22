package main

type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money{amount}}
}

func (d *Dollar) times(multiplier int) Dollar {
	return Dollar{Money{d.amount * multiplier}}
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

func (f *Franc) times(multiplier int) Franc {
	return Franc{Money{f.amount * multiplier}}
}

func (f *Franc) equals(anotherMoney any) bool {
	franc, ok := anotherMoney.(Franc)
	if !ok {
		return false
	}
	return f.amount == franc.amount
}

type Money struct {
	amount int
}

type MoneyI interface {
	equals(anotherMoney any) bool
}
