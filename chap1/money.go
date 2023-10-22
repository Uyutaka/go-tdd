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

type Franc struct {
	Money
}

func NewFranc(amount int) Franc {
	return Franc{Money{amount}}
}

func (f *Franc) times(multiplier int) Franc {
	return Franc{Money{f.amount * multiplier}}
}

type Money struct {
	amount int
}

func (m *Money) equals(anotherMoney any) bool {
	switch value := anotherMoney.(type) {
	case Dollar:
		return m.amount == value.amount
	case Franc:
		return m.amount == value.amount
	default:
		return false
	}
}
