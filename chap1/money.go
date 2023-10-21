package main

type Dollar struct {
	amount int
}

func New(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d *Dollar) times(multiplier int) {
	d.amount = multiplier * d.amount
}
