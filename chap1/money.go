package main

type Dollar struct {
	amount int
}

func New(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d *Dollar) times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}
