package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(iterable []Transaction, name string) float64 {
	var balance float64
	for _, t := range iterable {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
