package main

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d *Dollar) times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}

func (d *Dollar) equals(anotherDollar any) bool {
	return d.amount == anotherDollar.(Dollar).amount
}

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount: amount}
}

func (d *Franc) times(multiplier int) Franc {
	return Franc{amount: d.amount * multiplier}
}

func (f *Franc) equals(anotherFranc any) bool {
	return f.amount == anotherFranc.(Franc).amount
}
