package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(iterable []Transaction, name string) float64 {
	adjustBalance := func(currBalance float64, t Transaction) float64 {
		if t.From == name {
			return currBalance - t.Sum
		}
		if t.To == name {
			return currBalance + t.Sum
		}
		return currBalance
	}
	return Reduce(iterable, adjustBalance, 0.0)
}
